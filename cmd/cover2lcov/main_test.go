package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("converts coverage profile to lcov", func(t *testing.T) {
		tmpDir := t.TempDir()
		inPath := filepath.Join(tmpDir, "coverage.out")
		outPath := filepath.Join(tmpDir, "coverage.lcov")

		profile := strings.Join([]string{
			"mode: set",
			"pkg/a.go:10.2,10.10 1 1",
		}, "\n")

		if err := os.WriteFile(inPath, []byte(profile), 0o600); err != nil {
			t.Fatalf("WriteFile() error = %v", err)
		}

		if err := run(inPath, outPath); err != nil {
			t.Fatalf("run() error = %v", err)
		}

		outBytes, err := os.ReadFile(outPath)
		if err != nil {
			t.Fatalf("ReadFile() error = %v", err)
		}

		out := string(outBytes)
		for _, token := range []string{"TN:\n", "SF:pkg/a.go\n", "DA:10,1\n", "end_of_record\n"} {
			if !strings.Contains(out, token) {
				t.Fatalf("expected output to contain %q, got:\n%s", token, out)
			}
		}
	})

	t.Run("returns error when input is missing", func(t *testing.T) {
		tmpDir := t.TempDir()
		outPath := filepath.Join(tmpDir, "coverage.lcov")

		err := run(filepath.Join(tmpDir, "missing.out"), outPath)
		if err == nil {
			t.Fatal("expected error for missing input file")
		}
	})
}
