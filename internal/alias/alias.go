// Package alias maps fgit short codes and flaming-themed aliases to their
// corresponding git argument slices.
//
// Design notes:
//   - "gc" is intentionally absent: it conflicts with Go's garbage-collector
//     shorthand and several popular shell plugin sets (oh-my-zsh, etc.).
//   - "gl" maps to "git log" here; users of oh-my-zsh should prefer "flame"
//     since oh-my-zsh binds "gl" to "git pull".
//   - Compound commands (shipit, sync, review, amendit) are stored separately
//     and resolved via ExpandCompound to allow ordered multi-step execution.
package alias

// simple maps a single alias key to its git argument slice.
var simple = map[string][]string{
	// ── Status ───────────────────────────────────────────────────────────────
	"gs":  {"status"},
	"gss": {"status", "-s"},

	// ── Staging ──────────────────────────────────────────────────────────────
	"ga":  {"add"},
	"gaa": {"add", "-A"},
	"gap": {"add", "--patch"},
	"grs": {"restore"},

	// ── Commit ───────────────────────────────────────────────────────────────
	"gcm":  {"commit", "-m"},
	"gca":  {"commit", "--amend"},
	"gcan": {"commit", "--amend", "--no-edit"},

	// ── Branching ────────────────────────────────────────────────────────────
	"swap": {"switch"},
	"co":   {"checkout"},
	"cob":  {"checkout", "-b"},
	"gb":   {"branch"},
	"gba":  {"branch", "-a"},
	"gbd":  {"branch", "-d"},

	// ── History ──────────────────────────────────────────────────────────────
	"gl":  {"log"},
	"ggr": {"log", "--oneline", "--graph", "--decorate"},
	"gls": {"log", "--oneline", "-10"},

	// ── Diff ─────────────────────────────────────────────────────────────────
	"gd":  {"diff"},
	"gdc": {"diff", "--cached"},

	// ── Remote ───────────────────────────────────────────────────────────────
	"gp":   {"push"},
	"gps":  {"push", "-u", "origin", "HEAD"},
	"gpl":  {"pull"},
	"gf":   {"fetch"},
	"gfa":  {"fetch", "--all", "--prune"},
	"grv":  {"remote", "-v"},
	"gra":  {"remote", "add"},

	// ── Merge / Rebase ───────────────────────────────────────────────────────
	"gm":  {"merge"},
	"grb": {"rebase"},
	"gri": {"rebase", "--interactive"},

	// ── Cleanup ──────────────────────────────────────────────────────────────
	"burn":   {"reset", "--hard", "HEAD"},
	"gclean": {"clean", "-fd"},

	// ── Stash ────────────────────────────────────────────────────────────────
	"gst":  {"stash"},
	"gsp":  {"stash", "pop"},
	"gsl":  {"stash", "list"},
	"gstu": {"stash", "--include-untracked"},

	// ── Cherry-pick ──────────────────────────────────────────────────────────
	"gcp": {"cherry-pick"},

	// ── Tags ─────────────────────────────────────────────────────────────────
	"gt":  {"tag"},
	"gta": {"tag", "-a"},
	"gpt": {"push", "--tags"},

	// ── Bisect ───────────────────────────────────────────────────────────────
	"gbs": {"bisect", "start"},
	"gbb": {"bisect", "bad"},
	"gbg": {"bisect", "good"},

	// ── Misc ─────────────────────────────────────────────────────────────────
	"gbl":  {"blame"},
	"gref": {"reflog"},
	"gwt":  {"worktree", "add"},
	"gwtl": {"worktree", "list"},

	// ── Flaming word aliases ──────────────────────────────────────────────────
	"ignite":       {"add"},
	"ignite-all":   {"add", "-A"},
	"ignite-patch": {"add", "--patch"},
	"cool":         {"restore"},
	"blaze":        {"commit", "-m"},
	"rekindle":     {"commit", "--amend"},
	"rekindle-fast": {"commit", "--amend", "--no-edit"},
	"spawn":        {"checkout", "-b"},
	"flame":        {"log"},
	"inferno":      {"log", "--oneline", "--graph", "--decorate"},
	"spark":        {"log", "--oneline", "-10"},
	"scorch":       {"diff"},
	"scorch-staged": {"diff", "--cached"},
	"ship":         {"push"},
	"launch":       {"push", "-u", "origin", "HEAD"},
	"stoke":        {"pull"},
	"fuel":         {"fetch"},
	"refuel":       {"fetch", "--all", "--prune"},
	"fuse":         {"merge"},
	"twist":        {"rebase"},
	"twisty":       {"rebase", "--interactive"},
	"purge":        {"clean", "-fd"},
	"hide":         {"stash"},
	"unhide":       {"stash", "pop"},
	"ember":        {"cherry-pick"},
	"brand":        {"tag"},
	"hallmark":     {"tag", "-a"},
	"release":      {"push", "--tags"},
	"hunt":         {"bisect", "start"},
	"guilty":       {"bisect", "bad"},
	"innocent":     {"bisect", "good"},
	"inquisition":  {"blame"},
	"ashes":        {"reflog"},
	"campsite":     {"worktree", "add"},
	"campsites":    {"worktree", "list"},
}

// Step is a single git command expansion within a compound alias.
type Step struct {
	// Args are the git arguments for this step.
	Args []string
	// AppendUserArgs signals that any user-supplied extra arguments should be
	// appended to this step's Args. Only the last step in a compound uses this.
	AppendUserArgs bool
}

// compound maps multi-step aliases to their ordered git command sequences.
var compound = map[string][]Step{
	// shipit "message" → add -A, commit -m <message>, push
	"shipit": {
		{Args: []string{"add", "-A"}},
		{Args: []string{"commit", "-m"}, AppendUserArgs: true},
		{Args: []string{"push"}},
	},
	// sync → fetch, rebase
	"sync": {
		{Args: []string{"fetch"}},
		{Args: []string{"rebase"}},
	},
	// review → status, diff
	"review": {
		{Args: []string{"status"}},
		{Args: []string{"diff"}},
	},
	// amendit / rekindle-all → add -A, commit --amend --no-edit
	"amendit": {
		{Args: []string{"add", "-A"}},
		{Args: []string{"commit", "--amend", "--no-edit"}},
	},
	"rekindle-all": {
		{Args: []string{"add", "-A"}},
		{Args: []string{"commit", "--amend", "--no-edit"}},
	},
}

// Resolve returns the git argument slice for a simple alias.
// It returns (nil, false) when the alias is unknown or is a compound command.
func Resolve(key string) ([]string, bool) {
	args, ok := simple[key]
	return args, ok
}

// IsCompound reports whether key names a compound multi-step alias.
func IsCompound(key string) bool {
	_, ok := compound[key]
	return ok
}

// ExpandCompound returns the ordered Steps for a compound alias.
// It returns (nil, false) when key is not a compound alias.
func ExpandCompound(key string) ([]Step, bool) {
	steps, ok := compound[key]
	return steps, ok
}
