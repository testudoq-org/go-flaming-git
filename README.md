# 🔥 Flaming Git (`fgit`)

A lightweight, cross-platform Git wrapper that provides short 2-3 letter commands and fun flaming-themed aliases while fully supporting normal `git` usage via passthrough.

---

## Features

- Short codes (`gs`, `gcm`, `gp`) for the most common Git operations
- Flaming word aliases (`blaze`, `ship`, `inferno`) for personality and memorability
- Compound commands (`shipit`, `sync`, `review`) for multi-step workflows
- Full passthrough — any unknown command goes straight to `git`
- Single small binary with no runtime dependencies
- Works on Windows, Linux, and macOS

---

## Installation

### Download a pre-built binary

Download the appropriate binary for your platform from the [releases page](https://github.com/stephenhstewart/fgit/releases).

**Windows**
```powershell
# Place fgit-windows-amd64.exe somewhere on your PATH and rename it
Move-Item fgit-windows-amd64.exe C:\tools\fgit.exe
```

**Linux / macOS**
```bash
chmod +x fgit-linux-amd64
sudo mv fgit-linux-amd64 /usr/local/bin/fgit
```

### Build from source

Requirements: Go 1.22+

```bash
git clone https://github.com/stephenhstewart/fgit.git
cd fgit

# Quick local build
go build -o fgit .

# Cross-platform release build (all platforms)
./build.sh v1.0.0        # Linux/macOS
.\build.ps1 -Version v1.0.0  # Windows PowerShell
```

---

## Usage

```
fgit <alias|command> [args...]
```

All commands that are not known aliases pass through to `git` unchanged.

```bash
fgit gs                        # git status
fgit gss                       # git status -s
fgit gaa                       # git add -A
fgit blaze "Fix login bug"     # git commit -m "Fix login bug"
fgit ship                      # git push
fgit stoke                     # git pull
fgit swap main                 # git switch main
fgit spawn feature/new-thing   # git checkout -b feature/new-thing
fgit shipit "Released v1.0"    # git add -A && git commit -m "..." && git push
fgit sync                      # git fetch && git rebase
fgit inferno                   # git log --oneline --graph --decorate
fgit version                   # print fgit version
fgit aliases                   # print the full alias reference table
```

---

## Alias Reference

### Short Codes

| Alias   | Git Command                          | Notes                              |
|---------|--------------------------------------|------------------------------------|
| `gs`    | `git status`                         |                                    |
| `gss`   | `git status -s`                      |                                    |
| `ga`    | `git add`                            |                                    |
| `gaa`   | `git add -A`                         |                                    |
| `gap`   | `git add --patch`                    | Interactive staging                |
| `grs`   | `git restore`                        | Unstage / restore file             |
| `gcm`   | `git commit -m`                      |                                    |
| `gca`   | `git commit --amend`                 |                                    |
| `gcan`  | `git commit --amend --no-edit`       |                                    |
| `swap`  | `git switch`                         | Modern branch switching            |
| `co`    | `git checkout`                       | Legacy support                     |
| `cob`   | `git checkout -b`                    | Create and switch branch           |
| `gb`    | `git branch`                         |                                    |
| `gba`   | `git branch -a`                      |                                    |
| `gbd`   | `git branch -d`                      | Delete branch                      |
| `gl`    | `git log`                            | ⚠ Conflicts with oh-my-zsh        |
| `ggr`   | `git log --oneline --graph`          |                                    |
| `gls`   | `git log --oneline -10`              |                                    |
| `gd`    | `git diff`                           |                                    |
| `gdc`   | `git diff --cached`                  |                                    |
| `gp`    | `git push`                           |                                    |
| `gps`   | `git push -u origin HEAD`            | Set upstream                       |
| `gpl`   | `git pull`                           |                                    |
| `gf`    | `git fetch`                          |                                    |
| `gfa`   | `git fetch --all --prune`            |                                    |
| `gm`    | `git merge`                          |                                    |
| `grb`   | `git rebase`                         |                                    |
| `gri`   | `git rebase --interactive`           |                                    |
| `burn`  | `git reset --hard HEAD`              | ⚠ DANGEROUS                       |
| `gclean`| `git clean -fd`                      |                                    |
| `gst`   | `git stash`                          |                                    |
| `gsp`   | `git stash pop`                      |                                    |
| `gsl`   | `git stash list`                     |                                    |
| `gstu`  | `git stash --include-untracked`      |                                    |
| `gcp`   | `git cherry-pick`                    |                                    |
| `gt`    | `git tag`                            |                                    |
| `gta`   | `git tag -a`                         | Annotated tag                      |
| `gpt`   | `git push --tags`                    |                                    |
| `gbs`   | `git bisect start`                   |                                    |
| `gbb`   | `git bisect bad`                     |                                    |
| `gbg`   | `git bisect good`                    |                                    |
| `gbl`   | `git blame`                          |                                    |
| `gref`  | `git reflog`                         | Recover lost commits               |
| `gwt`   | `git worktree add`                   |                                    |
| `gwtl`  | `git worktree list`                  |                                    |

### Flaming Word Aliases

| Alias          | Git Command                          |
|----------------|--------------------------------------|
| `ignite`       | `git add`                            |
| `ignite-all`   | `git add -A`                         |
| `ignite-patch` | `git add --patch`                    |
| `cool`         | `git restore`                        |
| `blaze`        | `git commit -m`                      |
| `rekindle`     | `git commit --amend`                 |
| `rekindle-fast`| `git commit --amend --no-edit`       |
| `spawn`        | `git checkout -b`                    |
| `flame`        | `git log`                            |
| `inferno`      | `git log --oneline --graph`          |
| `spark`        | `git log --oneline -10`              |
| `scorch`       | `git diff`                           |
| `scorch-staged`| `git diff --cached`                  |
| `ship`         | `git push`                           |
| `launch`       | `git push -u origin HEAD`            |
| `stoke`        | `git pull`                           |
| `fuel`         | `git fetch`                          |
| `refuel`       | `git fetch --all --prune`            |
| `fuse`         | `git merge`                          |
| `twist`        | `git rebase`                         |
| `twisty`       | `git rebase --interactive`           |
| `purge`        | `git clean -fd`                      |
| `hide`         | `git stash`                          |
| `unhide`       | `git stash pop`                      |
| `ember`        | `git cherry-pick`                    |
| `brand`        | `git tag`                            |
| `hallmark`     | `git tag -a`                         |
| `release`      | `git push --tags`                    |
| `hunt`         | `git bisect start`                   |
| `guilty`       | `git bisect bad`                     |
| `innocent`     | `git bisect good`                    |
| `inquisition`  | `git blame`                          |
| `ashes`        | `git reflog`                         |
| `campsite`     | `git worktree add`                   |
| `campsites`    | `git worktree list`                  |

### Compound Commands

| Alias        | Expands To                                      | Notes                      |
|--------------|-------------------------------------------------|----------------------------|
| `shipit`     | `add -A` → `commit -m <msg>` → `push`          | Requires a commit message  |
| `sync`       | `fetch` → `rebase`                              |                            |
| `review`     | `status` → `diff`                              |                            |
| `amendit`    | `add -A` → `commit --amend --no-edit`           |                            |
| `rekindle-all`| `add -A` → `commit --amend --no-edit`          |                            |

---

## Warnings

> ⚠️ **`burn`** runs `git reset --hard HEAD`. All uncommitted changes are permanently lost.

> ⚠️ **`gl`** conflicts with oh-my-zsh which maps `gl` → `git pull`. Use `flame` instead if you use oh-my-zsh.

---

## Roadmap

| Version | Goals |
|---------|-------|
| v1 | Core aliases, compound commands, help system, cross-platform binary |
| v2 | User-configurable aliases via `~/.fgit.yaml` |
| v3 | Shell completion (bash, zsh, fish, PowerShell) |
| v4 | Interactive TUI mode for browsing aliases |

---

## Contributing

1. Follow the TDD workflow in `docs/tdd-workflow.md`
2. Run `go test ./...` before every commit
3. Ensure `golangci-lint run` passes
4. Keep changes small and focused — one behavior per commit

---

## License

MIT
