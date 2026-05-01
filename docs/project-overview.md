Here's a clean, well-structured **Overview Document** specifically designed as a **prompt guide** to build the complete Flaming Git (`fgit`) application.

This document breaks the entire project into **manageable, sequential chunks** so you can build it step-by-step without feeling overwhelmed.

---

### Flaming Git (`fgit`) – Build Overview & Prompt Guide

**Project Goal**:  
Create a lightweight, standalone, cross-platform Git wrapper called **Flaming Git** (`fgit`). It provides short 2-3 letter commands and fun flaming-themed aliases while fully supporting normal `git` usage via passthrough. The final output must be a single small executable that runs on Windows, Linux, and macOS.

**Core Constraints**:
- Must be a **true standalone binary** (no external runtime)
- Primary language: **Go**
- Keep binary size as small as possible
- Work reliably on Windows (primary dev), Linux (WSL), and macOS
- Use `fgit` as main binary name (`fg` as optional short alias)

---

### Overall Development Phases (Recommended Order)

I have broken the complete application into **8 manageable chunks**. Each chunk is self-contained and builds on the previous one.

#### Chunk 1: Project Setup & Basic Structure
**Goal**: Initialize the Go project and create the foundation.

**Key Tasks**:
- Initialize Go module
- Set up project folder structure
- Add Cobra CLI framework
- Create basic `main.go` with root command
- Implement simple passthrough to `git`

**Prompt Focus**: Project initialization, module setup, Cobra integration, basic command execution.

---

#### Chunk 2: Core Alias Mapping System
**Goal**: Build the heart of the application — the alias translation engine.

**Key Tasks**:
- Define expanded alias table (use the full table from previous response)
- Create internal mapping structure (`map[string]string`)
- Support both short codes (`gs`, `swap`) and flaming aliases (`ignite`, `shipit`)
- Handle compound commands (`shipit "msg"`, `sync`)
- Implement intelligent argument passing

**Prompt Focus**: Alias map definition, lookup logic, command expansion.

---

#### Chunk 3: Command Execution Engine
**Goal**: Safely execute Git commands and handle OS differences.

**Key Tasks**:
- Robust `runGit()` function
- Proper stdin/stdout/stderr forwarding
- Correct exit code handling
- Cross-platform path handling for `git` executable
- Error messaging

**Prompt Focus**: Process execution, I/O forwarding, error handling.

---

#### Chunk 4: Help System & User Experience
**Goal**: Make the tool discoverable and user-friendly.

**Key Tasks**:
- `fgit --help` and `fgit help`
- `fgit aliases` command with nicely formatted table
- `fgit h` as short for help
- Custom usage examples with flaming theme
- Warning notes (especially for `burn`, oh-my-zsh conflicts)

**Prompt Focus**: Cobra subcommands, formatted output, help text.

---

#### Chunk 5: Advanced Alias Features
**Goal**: Improve usability and flexibility.

**Key Tasks**:
- Support for compound commands with argument handling (e.g. `shipit "commit message"`)
- Smart parsing for commands that need special treatment
- Future-proof structure for config file (v2)
- Conflict-safe design enforcement

**Prompt Focus**: Argument parsing, compound command logic.

---

#### Chunk 6: Building & Cross-Compilation
**Goal**: Make it easy to build for all platforms.

**Key Tasks**:
- Create build scripts (PowerShell + Bash)
- Optimize binary size (`-ldflags="-s -w"`, `-trimpath`)
- Instructions for building Windows, Linux, and macOS binaries from Windows
- Add version information (`-ldflags` with build info)

**Prompt Focus**: Build optimization, cross-compilation, release preparation.

---

#### Chunk 7: Documentation & Distribution
**Goal**: Prepare the project for users.

**Key Tasks**:
- Professional `README.md`
- Installation instructions for Windows, Linux, macOS
- Full alias reference table (Markdown)
- Conflict audit summary
- Roadmap (v1 → v4)

**Prompt Focus**: Documentation writing, user guides.

---

#### Chunk 8: Polish, Testing & Refinement
**Goal**: Final quality improvements.

**Key Tasks**:
- Unit tests for alias mapping and command expansion
- Basic integration tests
- Input validation and safety checks
- Logging / verbose mode (optional)
- Final binary size optimization
- Code cleanup and comments

**Prompt Focus**: Testing, edge cases, optimization.

---

### How to Use This Document

You can copy any chunk below and use it as a **system + user prompt** when asking me (or another AI) to generate code.

---

### Detailed Chunk Prompts (Ready to Copy)

**Chunk 1 Prompt (Project Setup)**:
```
You are an expert Go developer. Create the initial project structure for "Flaming Git" (fgit) - a standalone Git wrapper.

Requirements:
- Go module name: github.com/stephenhstewart/fgit (or use a local name)
- Use spf13/cobra for CLI handling
- Main binary name: fgit
- Support running on Windows, Linux, macOS
- Create main.go with root command that currently just calls git with passthrough
- Include basic help text with flaming theme
- Project should compile to a single small binary

Provide the complete file structure and code for this first chunk.
```

**Chunk 2 Prompt (Core Alias Mapping)**:
```
Using the existing Flaming Git project, implement the full alias mapping system.

Include this complete alias table:
[ Paste the full Expanded Alias Table I gave you earlier ]

Requirements:
- Use a well-structured map or struct for aliases
- Support both short codes and flaming word aliases
- Handle compound commands like "shipit" and "sync"
- Add clear comments explaining design decisions (especially gc avoidance)
- Make the mapping easily extensible for future config file

Provide the updated code with the complete mapping implemented.
```

**Chunk 3 Prompt (Command Execution)**:
```
Enhance the Flaming Git project with a robust command execution engine.

Requirements:
- Create a function runGit(args []string) that properly executes git
- Forward stdin, stdout, stderr correctly
- Preserve original git exit codes
- Handle errors gracefully with clear messages
- Work reliably across Windows, Linux, and macOS

Update the code to use this execution layer.
```

**Chunk 4 Prompt (Help System)**:
```
Add a professional help system to Flaming Git.

Requirements:
- Nice `fgit --help` output
- `fgit aliases` command that displays all aliases in a clean table format
- Support `fgit h` as shortcut for help
- Include usage examples and warnings (burn command, oh-my-zsh gl conflict)
- Keep the flaming/fun personality in the help text

Provide the updated Cobra commands and help formatting code.
```
Notes on what to build

Here's the **expanded full alias table** for Flaming Git (`fgit`). I've made it comprehensive, practical, and conflict-aware based on common Git usage patterns and major alias sets (especially oh-my-zsh).

### 2. Expanded Alias Table (v1)

This table includes ~40 of the most useful Git operations. It prioritizes:
- **Safety** — avoiding heavily conflicted shorts like `gc` (used for commit in many setups).
- **Memorability** — 2-3 letter shorts + flaming-themed word aliases.
- **Modern Git** — preferring `switch` / `restore` where appropriate.
- **Compound commands** — for frequent multi-step workflows.

| Category              | Git Command                          | Short Code | Flaming Alias      | Notes / Conflicts |
|-----------------------|--------------------------------------|------------|--------------------|-------------------|
| **Status**            | status                               | gs         | -                  | Very safe |
| **Status**            | status -s                            | gss        | quick              | Short status |
| **Staging**           | add                                  | ga         | ignite             | Safe |
| **Staging**           | add -A / add .                       | gaa        | ignite-all         | Safe |
| **Staging**           | add --patch                          | gap        | ignite-patch       | - |
| **Staging**           | restore                              | grs        | cool               | Modern unstage |
| **Commit**            | commit -m                            | gcm        | blaze              | Preferred over `gc` |
| **Commit**            | commit                               | commit     | -                  | Passthrough |
| **Commit**            | commit --amend                       | gca        | rekindle           | - |
| **Commit**            | commit --amend --no-edit             | gcan       | rekindle-fast      | - |
| **Branching**         | switch                               | swap       | swap               | **Primary** for switching (modern) |
| **Branching**         | checkout                             | co         | -                  | Legacy support |
| **Branching**         | checkout -b                          | cob        | spawn              | Create + switch |
| **Branching**         | branch                               | gb         | branch             | - |
| **Branching**         | branch -a                            | gba        | branches-all       | - |
| **Branching**         | branch -d                            | gbd        | prune              | Delete branch |
| **History**           | log                                  | gl         | flame              | **Note**: oh-my-zsh maps `gl` → pull |
| **History**           | log --oneline --graph --decorate     | ggr        | inferno            | Renamed to avoid conflicts |
| **History**           | log --oneline -10                    | gls        | spark              | Short log |
| **Diff**              | diff                                 | gd         | scorch             | - |
| **Diff**              | diff --cached                        | gdc        | scorch-staged      | - |
| **Remote**            | push                                 | gp         | ship               | - |
| **Remote**            | push -u origin HEAD                  | gps        | launch             | Set upstream |
| **Remote**            | pull                                 | gpl        | stoke              | - |
| **Remote**            | fetch                                | gf         | fuel               | - |
| **Remote**            | fetch --all --prune                  | gfa        | refuel             | - |
| **Remote + Integrate**| fetch + rebase                       | sync       | sync               | Compound |
| **Merging**           | merge                                | gm         | fuse               | - |
| **Rebasing**          | rebase                               | grb        | twist              | - |
| **Rebasing**          | rebase --interactive                 | gri        | twisty             | - |
| **Cleanup**           | reset --hard HEAD                    | burn       | burn               | **Dangerous** – use with care |
| **Cleanup**           | clean -fd                            | gclean     | purge              | Remove untracked |
| **Stash**             | stash                                | gst        | hide               | Safe |
| **Stash**             | stash pop                            | gsp        | unhide             | - |
| **Stash**             | stash list                           | gsl        | hides              | - |
| **Compound**          | add -A + commit -m + push            | shipit     | shipit             | `shipit "message"` |
| **Compound**          | status + diff                        | review     | review             | Quick review |
| **Compound**          | add -A + commit --amend              | amendit    | rekindle-all       | - |
| **Setup**             | init                                 | init       | -                  | Passthrough |
| **Setup**             | clone                                | clone      | -                  | Passthrough |
| **Config**            | config --global                      | gconfig    | -                  | Passthrough |

### Flaming Theme Highlights
These word aliases give `fgit` its fun personality:
- **Ignite** family → staging (`ignite`, `ignite-all`)
- **Blaze** / **Rekindle** → committing
- **Swap** → branch switching (very intuitive)
- **Ship** / **Launch** → pushing
- **Stoke** → pulling (adding fuel to the fire)
- **Burn** → hard reset (destructive but memorable)
- **Scorch** → diff
- **Inferno** → pretty graph log

### How to Use
- Short codes: `fgit gs`, `fgit swap main`, `fgit gcm "Fixed bug"`
- Flaming aliases: `fgit shipit "Added login feature"`, `fgit burn`, `fgit ignite .`
- Passthrough: `fgit status --porcelain`, `fgit log --follow file.txt`

### Implementation Notes for Code
In the Go map, you can define compound commands like this:

```go
"shipit": "add -A && commit -m",
"sync":   "fetch && rebase",
"review": "status && diff",
```

For `shipit "message"`, the code needs to intelligently append the remaining arguments after the `-m`.

Would you like me to:
1. Provide an updated full Go source code with this expanded table + better compound command handling?
2. Add a "Suggested Aliases" section with even more flaming-themed options?
3. Include a Markdown version of this table ready for README.md?
4. Update the entire architecture document with this new expanded table?

Let me know your preference and any specific additions or tweaks (e.g., more stash commands, tag commands, or worktree support).