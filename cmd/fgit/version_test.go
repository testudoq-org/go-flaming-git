package fgit_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stephenhstewart/fgit/cmd/fgit"
)

func TestVersionCommand_IsRegistered(t *testing.T) {
	root := fgit.NewRootCommand()
	found := false
	for _, sub := range root.Commands() {
		if sub.Name() == "version" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected 'version' subcommand to be registered")
	}
}

func TestVersionCommand_OutputContainsVersion(t *testing.T) {
	root := fgit.NewRootCommand()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetArgs([]string{"version"})

	if err := root.Execute(); err != nil {
		t.Fatalf("version command failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "fgit") {
		t.Errorf("expected version output to contain 'fgit', got: %s", out)
	}
}
