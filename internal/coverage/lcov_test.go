package coverage

import (
	"bytes"
	"strings"
	"testing"
)

func TestConvertProfileToLCOV(t *testing.T) {
	input := strings.Join([]string{
		"mode: set",
		"pkg/a.go:10.2,12.4 1 1",
		"pkg/a.go:11.2,11.7 1 0",
		"pkg/b.go:3.1,3.8 1 2",
	}, "\n")

	var out bytes.Buffer
	err := ConvertProfileToLCOV(strings.NewReader(input), &out)
	if err != nil {
		t.Fatalf("ConvertProfileToLCOV() error = %v", err)
	}

	got := out.String()

	expectedLines := []string{
		"TN:",
		"SF:pkg/a.go",
		"DA:10,1",
		"DA:11,1",
		"DA:12,1",
		"end_of_record",
		"TN:",
		"SF:pkg/b.go",
		"DA:3,2",
		"end_of_record",
	}

	for _, line := range expectedLines {
		if !strings.Contains(got, line+"\n") {
			t.Fatalf("expected output to contain %q, got:\n%s", line, got)
		}
	}
}

func TestConvertProfileToLCOV_InvalidHeader(t *testing.T) {
	input := "pkg/a.go:10.2,12.4 1 1\n"

	var out bytes.Buffer
	err := ConvertProfileToLCOV(strings.NewReader(input), &out)
	if err == nil {
		t.Fatal("expected error for invalid header")
	}
}

func TestConvertProfileToLCOV_InvalidLine(t *testing.T) {
	input := strings.Join([]string{
		"mode: set",
		"pkg/a.go:10.2,12.4 1 not-a-number",
	}, "\n")

	var out bytes.Buffer
	err := ConvertProfileToLCOV(strings.NewReader(input), &out)
	if err == nil {
		t.Fatal("expected parse error for invalid execution count")
	}
}

func TestCollectLineCoverage(t *testing.T) {
	files := make(fileLines)
	err := collectLineCoverage(files, "pkg/a.go:8.1,10.2 1 2")
	if err != nil {
		t.Fatalf("collectLineCoverage() error = %v", err)
	}

	for _, lineNo := range []int{8, 9, 10} {
		if got := files["pkg/a.go"][lineNo]; got != 2 {
			t.Fatalf("line %d count = %d, want 2", lineNo, got)
		}
	}
}

func TestParseLineNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid", input: "12.34", want: 12},
		{name: "missing dot", input: "12", wantErr: true},
		{name: "invalid number", input: "a.1", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLineNumber(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("parseLineNumber(%q) expected error", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("parseLineNumber(%q) unexpected error: %v", tt.input, err)
			}

			if got != tt.want {
				t.Fatalf("parseLineNumber(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
