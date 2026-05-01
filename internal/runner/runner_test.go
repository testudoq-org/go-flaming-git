package runner_test

import (
	"testing"

	"github.com/stephenhstewart/fgit/internal/runner"
)

func TestRun_ReturnsErrorOnMissingExecutable(t *testing.T) {
	err := runner.Run("__nonexistent_binary__", []string{"--version"})
	if err == nil {
		t.Error("expected error for missing executable, got nil")
	}
}

func TestRun_GitVersionSucceeds(t *testing.T) {
	err := runner.Run("git", []string{"--version"})
	if err != nil {
		t.Errorf("expected git --version to succeed, got: %v", err)
	}
}
