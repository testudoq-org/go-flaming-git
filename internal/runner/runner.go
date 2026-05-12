// Package runner executes external commands, forwarding stdin/stdout/stderr
// and preserving the original exit code.
package runner

import (
	"errors"
	"os"
	"os/exec"
)

// Run executes the named binary with the provided arguments.
// stdin, stdout, and stderr are forwarded to the parent process.
// An error is returned when the binary cannot be found or the process exits
// with a non-zero status.
func Run(binary string, args []string) error {
	path, err := exec.LookPath(binary)
	if err != nil {
		return errors.New("fgit: command not found: " + binary)
	}

	cmd := exec.Command(path, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
