// Package fgit defines the Cobra root command for the fgit binary.
package fgit

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/stephenhstewart/fgit/internal/dispatch"
	"github.com/stephenhstewart/fgit/internal/runner"
)

// NewRootCommand builds and returns the root Cobra command.
func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "fgit",
		Short: "Flaming Git — fast aliases and flaming-themed shortcuts for git",
		Long: `
 🔥 Flaming Git (fgit)

 A lightweight, cross-platform Git wrapper with short aliases
 and flaming-themed commands. All unknown commands pass straight
 through to git.

 Examples:
   fgit gs              # git status
   fgit blaze "msg"     # git commit -m "msg"
   fgit shipit "msg"    # git add -A && git commit -m "msg" && git push

 Run 'fgit aliases' for the full alias reference.`,
		// Disable default completion command
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		// DisableFlagParsing passes all args to RunE so git flags like --oneline
		// are not intercepted by Cobra. Subcommands (aliases, version) are still
		// matched by Cobra before RunE is reached.
		DisableFlagParsing: true,
		Args:               cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			d := dispatch.New(func(gitArgs []string) error {
				return runner.Run("git", gitArgs)
			})
			return d.Dispatch(args)
		},
		// Silence usage output on error; let git handle its own messages
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(newAliasesCommand())
	root.AddCommand(newVersionCommand())
	return root
}

// Execute runs the root command. It calls os.Exit on failure.
func Execute() {
	root := NewRootCommand()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
