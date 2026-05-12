package fgit

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// aliasRow describes one row in the aliases reference table.
type aliasRow struct {
	shortCode   string
	flamingWord string
	gitCommand  string
	notes       string
}

// allAliasRows is the full reference table printed by `fgit aliases`.
var allAliasRows = []aliasRow{
	// Status
	{"gs", "-", "git status", ""},
	{"gss", "quick", "git status -s", "Short status"},
	// Staging
	{"ga", "ignite", "git add", ""},
	{"gaa", "ignite-all", "git add -A", ""},
	{"gap", "ignite-patch", "git add --patch", "Interactive staging"},
	{"grs", "cool", "git restore", "Modern unstage"},
	// Commit
	{"gcm", "blaze", "git commit -m", ""},
	{"gca", "rekindle", "git commit --amend", ""},
	{"gcan", "rekindle-fast", "git commit --amend --no-edit", ""},
	// Branching
	{"swap", "swap", "git switch", "Modern branch switch"},
	{"co", "-", "git checkout", "Legacy support"},
	{"cob", "spawn", "git checkout -b", "Create + switch"},
	{"gb", "-", "git branch", ""},
	{"gba", "-", "git branch -a", ""},
	{"gbd", "prune", "git branch -d", "Delete branch"},
	// History
	{"gl", "flame", "git log", "⚠ oh-my-zsh maps gl→pull; use 'flame' instead"},
	{"ggr", "inferno", "git log --oneline --graph --decorate", ""},
	{"gls", "spark", "git log --oneline -10", ""},
	// Diff
	{"gd", "scorch", "git diff", ""},
	{"gdc", "scorch-staged", "git diff --cached", ""},
	// Remote
	{"gp", "ship", "git push", ""},
	{"gps", "launch", "git push -u origin HEAD", "Set upstream"},
	{"gpl", "stoke", "git pull", ""},
	{"gf", "fuel", "git fetch", ""},
	{"gfa", "refuel", "git fetch --all --prune", ""},
	{"grv", "-", "git remote -v", ""},
	{"gra", "-", "git remote add", ""},
	// Merge / Rebase
	{"gm", "fuse", "git merge", ""},
	{"grb", "twist", "git rebase", ""},
	{"gri", "twisty", "git rebase --interactive", ""},
	// Cleanup
	{"burn", "burn", "git reset --hard HEAD", "⚠ DANGEROUS — discards all local changes"},
	{"gclean", "purge", "git clean -fd", "Remove untracked files"},
	// Stash
	{"gst", "hide", "git stash", ""},
	{"gsp", "unhide", "git stash pop", ""},
	{"gsl", "-", "git stash list", ""},
	{"gstu", "-", "git stash --include-untracked", ""},
	// Cherry-pick
	{"gcp", "ember", "git cherry-pick", ""},
	// Tags
	{"gt", "brand", "git tag", ""},
	{"gta", "hallmark", "git tag -a", "Annotated tag"},
	{"gpt", "release", "git push --tags", "Push all tags"},
	// Bisect
	{"gbs", "hunt", "git bisect start", ""},
	{"gbb", "guilty", "git bisect bad", ""},
	{"gbg", "innocent", "git bisect good", ""},
	// Misc
	{"gbl", "inquisition", "git blame", ""},
	{"gref", "ashes", "git reflog", "Recover lost commits"},
	{"gwt", "campsite", "git worktree add", ""},
	{"gwtl", "campsites", "git worktree list", ""},
	// Compound
	{"shipit", "shipit", "add -A && commit -m && push", `shipit "message"`},
	{"sync", "sync", "fetch && rebase", ""},
	{"review", "review", "status && diff", ""},
	{"amendit", "rekindle-all", "add -A && commit --amend --no-edit", ""},
}

// newAliasesCommand builds the `fgit aliases` subcommand.
func newAliasesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "aliases",
		Short: "List all fgit aliases and their git equivalents",
		Long:  "Display the full alias reference table including short codes, flaming word aliases, and compound commands.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			printAliasTable(cmd.OutOrStdout())
			return nil
		},
	}
}

func printAliasTable(w io.Writer) {
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, " 🔥 Flaming Git — Alias Reference")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "  %-14s  %-20s  %-42s  %s\n", "Short Code", "Alias", "Git Command", "Notes")
	fmt.Fprintf(w, "  %-14s  %-20s  %-42s  %s\n",
		"--------------", "--------------------", "------------------------------------------", "-----")
	for _, row := range allAliasRows {
		fmt.Fprintf(w, "  %-14s  %-20s  %-42s  %s\n",
			row.shortCode, row.flamingWord, row.gitCommand, row.notes)
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  WARNING: 'burn' resets --hard HEAD. All uncommitted changes are lost.")
	fmt.Fprintln(w, "  WARNING: 'gl' conflicts with oh-my-zsh (which maps gl→pull). Use 'flame' instead.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  Usage examples:")
	fmt.Fprintln(w, `    fgit gs                    # git status`)
	fmt.Fprintln(w, `    fgit blaze "Fix login bug"  # git commit -m "Fix login bug"`)
	fmt.Fprintln(w, `    fgit shipit "Released v1"   # add -A, commit -m, push`)
	fmt.Fprintln(w, `    fgit swap main              # git switch main`)
	fmt.Fprintln(w, "")
}
