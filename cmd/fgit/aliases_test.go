package fgit_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stephenhstewart/fgit/cmd/fgit"
)

func TestAliasesCommand_IsRegistered(t *testing.T) {
	root := fgit.NewRootCommand()
	found := false
	for _, sub := range root.Commands() {
		if sub.Name() == "aliases" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected 'aliases' subcommand to be registered on root")
	}
}

func TestAliasesCommand_OutputContainsHeaders(t *testing.T) {
	root := fgit.NewRootCommand()
	buf := &bytes.Buffer{}
	root.SetOut(buf)

	root.SetArgs([]string{"aliases"})
	if err := root.Execute(); err != nil {
		t.Fatalf("aliases command failed: %v", err)
	}

	out := buf.String()
	for _, want := range []string{"Short Code", "Git Command", "Alias"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected output to contain %q, got:\n%s", want, out)
		}
	}
}

func TestAliasesCommand_OutputContainsSampleAliases(t *testing.T) {
	root := fgit.NewRootCommand()
	buf := &bytes.Buffer{}
	root.SetOut(buf)

	root.SetArgs([]string{"aliases"})
	if err := root.Execute(); err != nil {
		t.Fatalf("aliases command failed: %v", err)
	}

	out := buf.String()
	for _, want := range []string{"shipit", "blaze", "burn", "gs", "swap"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected output to contain alias %q", want)
		}
	}
}

func TestAliasesCommand_OutputContainsWarning(t *testing.T) {
	root := fgit.NewRootCommand()
	buf := &bytes.Buffer{}
	root.SetOut(buf)

	root.SetArgs([]string{"aliases"})
	if err := root.Execute(); err != nil {
		t.Fatalf("aliases command failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "WARNING") && !strings.Contains(out, "warning") {
		t.Error("expected output to contain a warning about dangerous commands or conflicts")
	}
}
