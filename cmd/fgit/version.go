package fgit

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is injected at build time via -ldflags.
// Default value is used when building without ldflags.
var Version = "dev"

// newVersionCommand builds the `fgit version` subcommand.
func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the fgit version",
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "fgit version %s\n", Version)
			return nil
		},
	}
}
