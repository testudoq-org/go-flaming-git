package alias_test

import (
	"testing"

	"github.com/stephenhstewart/fgit/internal/alias"
)

func TestResolve_KnownShortCode(t *testing.T) {
	tests := []struct {
		input    string
		wantArgs []string
	}{
		{"gs", []string{"status"}},
		{"gss", []string{"status", "-s"}},
		{"ga", []string{"add"}},
		{"gaa", []string{"add", "-A"}},
		{"gap", []string{"add", "--patch"}},
		{"grs", []string{"restore"}},
		{"gcm", []string{"commit", "-m"}},
		{"gca", []string{"commit", "--amend"}},
		{"gcan", []string{"commit", "--amend", "--no-edit"}},
		{"swap", []string{"switch"}},
		{"co", []string{"checkout"}},
		{"cob", []string{"checkout", "-b"}},
		{"gb", []string{"branch"}},
		{"gba", []string{"branch", "-a"}},
		{"gbd", []string{"branch", "-d"}},
		{"gl", []string{"log"}},
		{"ggr", []string{"log", "--oneline", "--graph", "--decorate"}},
		{"gls", []string{"log", "--oneline", "-10"}},
		{"gd", []string{"diff"}},
		{"gdc", []string{"diff", "--cached"}},
		{"gp", []string{"push"}},
		{"gps", []string{"push", "-u", "origin", "HEAD"}},
		{"gpl", []string{"pull"}},
		{"gf", []string{"fetch"}},
		{"gfa", []string{"fetch", "--all", "--prune"}},
		{"gm", []string{"merge"}},
		{"grb", []string{"rebase"}},
		{"gri", []string{"rebase", "--interactive"}},
		{"burn", []string{"reset", "--hard", "HEAD"}},
		{"gclean", []string{"clean", "-fd"}},
		{"gst", []string{"stash"}},
		{"gsp", []string{"stash", "pop"}},
		{"gsl", []string{"stash", "list"}},
		{"gstu", []string{"stash", "--include-untracked"}},
		{"grv", []string{"remote", "-v"}},
		{"gra", []string{"remote", "add"}},
		{"gcp", []string{"cherry-pick"}},
		{"gt", []string{"tag"}},
		{"gta", []string{"tag", "-a"}},
		{"gpt", []string{"push", "--tags"}},
		{"gbs", []string{"bisect", "start"}},
		{"gbb", []string{"bisect", "bad"}},
		{"gbg", []string{"bisect", "good"}},
		{"gbl", []string{"blame"}},
		{"gref", []string{"reflog"}},
		{"gwt", []string{"worktree", "add"}},
		{"gwtl", []string{"worktree", "list"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, ok := alias.Resolve(tt.input)
			if !ok {
				t.Fatalf("Resolve(%q) not found", tt.input)
			}
			if len(got) != len(tt.wantArgs) {
				t.Fatalf("Resolve(%q) = %v, want %v", tt.input, got, tt.wantArgs)
			}
			for i := range got {
				if got[i] != tt.wantArgs[i] {
					t.Errorf("Resolve(%q)[%d] = %q, want %q", tt.input, i, got[i], tt.wantArgs[i])
				}
			}
		})
	}
}

func TestResolve_FlamingWordAliases(t *testing.T) {
	tests := []struct {
		input    string
		wantArgs []string
	}{
		{"ignite", []string{"add"}},
		{"ignite-all", []string{"add", "-A"}},
		{"ignite-patch", []string{"add", "--patch"}},
		{"cool", []string{"restore"}},
		{"blaze", []string{"commit", "-m"}},
		{"rekindle", []string{"commit", "--amend"}},
		{"rekindle-fast", []string{"commit", "--amend", "--no-edit"}},
		{"spawn", []string{"checkout", "-b"}},
		{"flame", []string{"log"}},
		{"inferno", []string{"log", "--oneline", "--graph", "--decorate"}},
		{"spark", []string{"log", "--oneline", "-10"}},
		{"scorch", []string{"diff"}},
		{"scorch-staged", []string{"diff", "--cached"}},
		{"ship", []string{"push"}},
		{"launch", []string{"push", "-u", "origin", "HEAD"}},
		{"stoke", []string{"pull"}},
		{"fuel", []string{"fetch"}},
		{"refuel", []string{"fetch", "--all", "--prune"}},
		{"fuse", []string{"merge"}},
		{"twist", []string{"rebase"}},
		{"twisty", []string{"rebase", "--interactive"}},
		{"purge", []string{"clean", "-fd"}},
		{"hide", []string{"stash"}},
		{"unhide", []string{"stash", "pop"}},
		{"ember", []string{"cherry-pick"}},
		{"brand", []string{"tag"}},
		{"hallmark", []string{"tag", "-a"}},
		{"release", []string{"push", "--tags"}},
		{"hunt", []string{"bisect", "start"}},
		{"guilty", []string{"bisect", "bad"}},
		{"innocent", []string{"bisect", "good"}},
		{"inquisition", []string{"blame"}},
		{"ashes", []string{"reflog"}},
		{"campsite", []string{"worktree", "add"}},
		{"campsites", []string{"worktree", "list"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, ok := alias.Resolve(tt.input)
			if !ok {
				t.Fatalf("Resolve(%q) not found", tt.input)
			}
			if len(got) != len(tt.wantArgs) {
				t.Fatalf("Resolve(%q) = %v, want %v", tt.input, got, tt.wantArgs)
			}
			for i := range got {
				if got[i] != tt.wantArgs[i] {
					t.Errorf("Resolve(%q)[%d] = %q, want %q", tt.input, i, got[i], tt.wantArgs[i])
				}
			}
		})
	}
}

func TestResolve_UnknownAliasReturnsFalse(t *testing.T) {
	_, ok := alias.Resolve("totally-unknown-command")
	if ok {
		t.Error("expected Resolve to return false for unknown alias")
	}
}

func TestIsCompound_TrueForCompoundAliases(t *testing.T) {
	compounds := []string{"shipit", "sync", "review", "amendit", "rekindle-all"}
	for _, c := range compounds {
		t.Run(c, func(t *testing.T) {
			if !alias.IsCompound(c) {
				t.Errorf("expected %q to be compound", c)
			}
		})
	}
}

func TestIsCompound_FalseForSimpleAliases(t *testing.T) {
	if alias.IsCompound("gs") {
		t.Error("expected gs to not be compound")
	}
}

func TestExpandCompound_Shipit(t *testing.T) {
	steps, ok := alias.ExpandCompound("shipit")
	if !ok {
		t.Fatal("expected shipit to be a compound command")
	}
	if len(steps) != 3 {
		t.Fatalf("expected 3 steps for shipit, got %d", len(steps))
	}
}

func TestExpandCompound_Sync(t *testing.T) {
	steps, ok := alias.ExpandCompound("sync")
	if !ok {
		t.Fatal("expected sync to be a compound command")
	}
	if len(steps) != 2 {
		t.Fatalf("expected 2 steps for sync, got %d", len(steps))
	}
}
