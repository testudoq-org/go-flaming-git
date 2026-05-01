package fgit_test

import (
	"testing"

	"github.com/stephenhstewart/fgit/cmd/fgit"
)

func TestNewRootCommand_NotNil(t *testing.T) {
	cmd := fgit.NewRootCommand()
	if cmd == nil {
		t.Error("expected root command, got nil")
	}
}

func TestNewRootCommand_HasCorrectUse(t *testing.T) {
	cmd := fgit.NewRootCommand()
	if cmd.Use != "fgit" {
		t.Errorf("expected Use=fgit, got %q", cmd.Use)
	}
}
