// Package dispatch resolves a fgit command (alias, compound, or passthrough)
// and executes the corresponding git argument sequences via a RunFunc.
package dispatch

import (
	"errors"

	"github.com/stephenhstewart/fgit/internal/alias"
)

// requiresMessage lists aliases whose git expansion ends with -m, meaning the
// user must supply a commit message as the next argument.
var requiresMessage = map[string]bool{
	"gcm":    true,
	"blaze":  true,
	"shipit": true,
}

// RunFunc is the signature of the function that executes a single git invocation.
type RunFunc func(args []string) error

// Dispatcher resolves and runs fgit commands.
type Dispatcher struct {
	run RunFunc
}

// New creates a Dispatcher that delegates execution to run.
func New(run RunFunc) *Dispatcher {
	return &Dispatcher{run: run}
}

// Dispatch resolves args[0] as a fgit alias (simple or compound) or
// passes the full args slice directly to git as a passthrough.
func (d *Dispatcher) Dispatch(args []string) error {
	if len(args) == 0 {
		return d.run(args)
	}

	key := args[0]
	extra := args[1:]

	// Validate that commit-message aliases are called with a message.
	if requiresMessage[key] && len(extra) == 0 {
		return errors.New("fgit: '" + key + "' requires a commit message, e.g. fgit " + key + " \"your message\"")
	}

	// Compound alias: multiple sequential git calls.
	if alias.IsCompound(key) {
		steps, _ := alias.ExpandCompound(key)
		for _, step := range steps {
			stepArgs := append([]string(nil), step.Args...)
			if step.AppendUserArgs {
				stepArgs = append(stepArgs, extra...)
			}
			if err := d.run(stepArgs); err != nil {
				return err
			}
		}
		return nil
	}

	// Simple alias: single git call with optional extra args appended.
	if resolved, ok := alias.Resolve(key); ok {
		combined := append(append([]string(nil), resolved...), extra...)
		return d.run(combined)
	}

	// Passthrough: not a known alias, send args straight to git.
	return d.run(args)
}
