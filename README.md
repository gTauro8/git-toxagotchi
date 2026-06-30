# 🐉 Git-Toxagotchi

```
   .---.
  /     \
 | o   o |
  \  ~  /
   '---'
  Git-Toxagotchi
```

> A Tamagotchi-style ASCII pet that lives in your terminal and evolves based on the quality of your Git habits.

[![CI](https://github.com/sugar_petauro/git-toxagotchi/actions/workflows/ci.yml/badge.svg)](https://github.com/sugar_petauro/git-toxagotchi/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sugar_petauro/git-toxagotchi)](https://goreportcard.com/report/github.com/sugar_petauro/git-toxagotchi)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<!-- GIF placeholder: demo.gif -->

---

## What is this?

Git-Toxagotchi watches your Git habits and reacts accordingly. Write clean commit messages? Your pet evolves. Force push to main? It gets sarcastic. Add 47 new dependencies? It questions your life choices.

It's harmless, funny, and oddly motivating.

**It never modifies your files, commits, or Git history. Ever.**

---

## Install

```bash
go install github.com/sugar_petauro/git-toxagotchi/cmd/git-toxagotchi@latest
```

Or download a pre-built binary from the [releases page](https://github.com/sugar_petauro/git-toxagotchi/releases).

---

## Quickstart

```bash
# Create your pet
git-toxagotchi init --name "Patch"

# Check its status
git-toxagotchi status

# Analyze your staged changes before committing
git-toxagotchi analyze

# Open the interactive TUI
git-toxagotchi watch

# Install the pre-commit hook (opt-in)
git-toxagotchi hook install

# See your achievements
git-toxagotchi achievements

# Generate a badge for your README
git-toxagotchi share
```

---

## Commands

| Command | Description |
|---|---|
| `init` | Create a new pet |
| `status` | Show pet stats |
| `watch` | Open interactive TUI |
| `feed` | Feed your pet |
| `play` | Play with your pet |
| `analyze` | Analyze staged changes |
| `hook install` | Install pre-commit hook (opt-in) |
| `hook uninstall` | Remove hook |
| `achievements` | List achievements |
| `share` | Generate README badge |
| `config` | Show/edit configuration |

---

## TUI Controls

| Key | Action |
|---|---|
| `f` | Feed |
| `p` | Play |
| `s` | Status |
| `h` | Help |
| `q` | Quit |

---

## Evolution Stages

Your pet evolves as you accumulate experience through good commits:

```
Egg → Blob → Gremlin → Developer Cat → Terminal Wizard → Kernel Penguin → Git Dragon → Legendary Maintainer
```

---

## Achievements

- 🐣 **First Commit** — Made your first commit
- ✨ **10 Good Commits** — 10 commits with quality messages
- 🧪 **100 Tests Passed** — 100 tests passed
- 🕊️ **No Force Push Week** — A week without force pushing
- 🔍 **Secret Hunter** — Caught a secret before it was committed
- 📖 **Documentation Enjoyer** — Updated docs 5 times
- 🥗 **Dependency Diet** — No new dependencies for a week
- ⚔️ **Merge Conflict Survivor** — Survived a merge conflict

---

## Safety Principles

Git-Toxagotchi is **read-only by design**:

- ❌ Never modifies files, commits, or history
- ❌ Never writes to remotes
- ❌ No destructive filesystem operations
- ❌ No network calls in the core
- ✅ Pre-commit hook is strictly opt-in
- ✅ Hook only blocks commits containing secrets
- ✅ All "chaos" is simulated text — no actual chaos
- ✅ Works fully offline

---

## Configuration

Config lives at `~/.config/git-toxagotchi/config.yaml`:

```yaml
pet_name: "Patch"
theme: "default"
llm_enabled: false
hook_blocking: true
db_path: "~/.local/share/git-toxagotchi/pets.db"
```

---

## Architecture

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).

---

## Plugin System

Git-Toxagotchi supports plugins. See [docs/PLUGIN_GUIDE.md](docs/PLUGIN_GUIDE.md).

---

## Roadmap

See [docs/ROADMAP.md](docs/ROADMAP.md).

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). We welcome PRs!

---

## License

MIT — see [LICENSE](LICENSE).
