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

[![CI](https://github.com/gTauro8/git-toxagotchi/actions/workflows/ci.yml/badge.svg)](https://github.com/gTauro8/git-toxagotchi/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gTauro8/git-toxagotchi)](https://goreportcard.com/report/github.com/gTauro8/git-toxagotchi)
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
go install github.com/gTauro8/git-toxagotchi/cmd/git-toxagotchi@latest
```

Or download a pre-built binary from the [releases page](https://github.com/gTauro8/git-toxagotchi/releases).

---

## Quickstart

```bash
# First launch — opens an interactive setup wizard
git-toxagotchi init
```

The wizard walks you through:
1. **Pet name** — what to call your companion
2. **Theme** — choose between `default`, `dracula`, `nord`, `solarized`
3. **Hook blocking** — whether the pre-commit hook should block dangerous commits

After setup:

```bash
# Check your pet's status
git-toxagotchi status

# Analyze staged changes before committing
git-toxagotchi analyze

# Open the interactive TUI
git-toxagotchi watch

# Install the pre-commit hook (opt-in)
git-toxagotchi hook install

# See your achievements
git-toxagotchi achievements

# Generate a README badge
git-toxagotchi share
```

---

## Commands

| Command | Description |
|---|---|
| `init` | Create a new pet (runs setup wizard on first launch) |
| `status` | Show pet stats |
| `watch` | Open interactive TUI |
| `feed` | Feed your pet |
| `play` | Play with your pet |
| `analyze` | Analyze staged changes |
| `hook install` | Install pre-commit hook (opt-in) |
| `hook uninstall` | Remove hook |
| `achievements` | List achievements |
| `share` | Generate README badge |
| `config show` | Show current configuration |
| `config set <key> <value>` | Edit a single config field |
| `config edit` | Re-open the setup wizard |

---

## Configuration

On first launch, `git-toxagotchi init` opens an interactive wizard to configure your pet. You can re-run it any time with `config edit`, or change individual values with `config set`:

```bash
git-toxagotchi config set pet_name "Kernel"
git-toxagotchi config set theme dracula
git-toxagotchi config set hook_blocking true
```

Available keys: `pet_name`, `theme`, `llm_enabled`, `hook_blocking`

Config is saved to `~/.config/git-toxagotchi/config.yaml`:

```yaml
pet_name: "Patch"
theme: "default"
llm_enabled: false
hook_blocking: true
```

To skip the wizard in scripts or CI:

```bash
git-toxagotchi init --no-wizard --name "Toxi"
```

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
