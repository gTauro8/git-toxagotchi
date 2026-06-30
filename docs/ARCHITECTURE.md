# Architecture

Git-Toxagotchi follows a clean layered architecture.

## Layers

```
cmd/git-toxagotchi/     — CLI entry point (Cobra)
internal/
  domain/               — Core business logic: Pet, Events, Achievements, Evolution
  application/          — Use cases: Service, Analyzer, HumorEngine
  infrastructure/
    git/                — Read-only Git shell integration
    storage/            — SQLite persistence
    config/             — YAML config loading
  tui/                  — Bubble Tea TUI
  plugins/              — Plugin interface + built-in plugins
assets/                 — Embedded ASCII art
```

## Data Flow

```
User action
  → CLI command (Cobra)
    → Application Service
      → Git Reader (read-only)
      → Analyzer (commit analysis)
      → Domain (pet state update)
      → Storage (SQLite persist)
      → HumorEngine (response text)
    → TUI or stdout
```

## Key Design Decisions

**No CGO**: Uses `modernc.org/sqlite` (pure Go SQLite) so binaries are fully static.

**Read-only Git**: All git operations use `exec.Command("git", ...)` with read-only commands (`diff`, `log`, `status`). The binary never calls `git commit`, `git push`, or anything destructive.

**Embedded assets**: ASCII art is embedded at compile time via `//go:embed`, so the binary is self-contained.

**Plugin system**: Plugins implement a simple interface and receive events. They cannot modify pet state directly — they observe.

**Offline-first**: No network calls in the core. The `share` command generates local files only.

## Storage Schema

```sql
CREATE TABLE pets (
  id TEXT PRIMARY KEY,
  name TEXT,
  species TEXT,
  stage TEXT,
  energy INTEGER,
  hunger INTEGER,
  stress INTEGER,
  trust INTEGER,
  chaos INTEGER,
  experience INTEGER,
  age INTEGER,
  mood TEXT,
  last_interaction_at DATETIME
);

CREATE TABLE events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  pet_id TEXT,
  type TEXT,
  metadata TEXT,
  created_at DATETIME
);

CREATE TABLE achievements (
  id TEXT PRIMARY KEY,
  pet_id TEXT,
  unlocked INTEGER,
  unlocked_at DATETIME
);
```
