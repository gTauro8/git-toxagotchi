# Roadmap

## v0.1 — MVP (current)

- [x] Core pet domain model (mood, energy, stress, trust, chaos)
- [x] 8 evolution stages
- [x] 8 achievements
- [x] Commit analyzer (message quality, secrets, TODOs, diff size)
- [x] Humor engine with 50+ sarcastic lines
- [x] Fake chaos messages
- [x] Bubble Tea TUI with ASCII art
- [x] SQLite persistence
- [x] YAML config
- [x] Pre-commit hook (opt-in)
- [x] Share command (local badge + markdown)
- [x] Plugin architecture

## v0.2 — Polish

- [ ] More ASCII art variants per mood
- [ ] Sound effects (terminal bell)
- [ ] `config set` subcommand
- [ ] Pet naming via TUI prompt
- [ ] `git-toxagotchi reset` to start over
- [ ] Better secret detection (more patterns)
- [ ] Test result integration (`go test` output parsing)

## v0.3 — Social

- [ ] GitHub Gist badge upload (opt-in, explicit consent)
- [ ] Export pet as JSON for sharing
- [ ] Import friend's pet stats (read-only leaderboard)
- [ ] GitHub Actions workflow integration

## v0.4 — LLM (Optional)

- [ ] Optional LLM backend for dynamic commit feedback
- [ ] Configurable: local Ollama, Claude API, or OpenAI
- [ ] LLM strictly opt-in, disabled by default
- [ ] LLM prompt safety guardrails

## v1.0 — Stable

- [ ] Stable plugin API
- [ ] Homebrew formula
- [ ] Windows installer
- [ ] i18n support
- [ ] Documentation site
