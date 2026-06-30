# Contributing to Git-Toxagotchi

Thank you for your interest in contributing!

## Getting Started

```bash
git clone https://github.com/sugar_petauro/git-toxagotchi
cd git-toxagotchi
go mod download
go test ./...
```

## Development Guidelines

- Keep functions small and focused
- Write tests for new behavior
- Run `gofmt -w .` before committing
- Run `golangci-lint run` before opening a PR
- No CGO — use pure Go libraries

## Pull Request Process

1. Fork the repo
2. Create a feature branch: `git checkout -b feat/your-feature`
3. Write tests for your changes
4. Ensure `go test ./...` passes
5. Open a PR with a clear description

## Adding Humor Lines

The humor engine lives in `internal/application/humor.go`. Add new lines to the appropriate category slice. Keep the tone funny, not mean.

## Adding Achievements

1. Add a new constant to `internal/domain/achievements.go`
2. Add it to the `AllAchievements()` slice
3. Add unlock logic in `internal/application/service.go`
4. Write a test

## Adding Plugins

See [docs/PLUGIN_GUIDE.md](docs/PLUGIN_GUIDE.md).

## Code of Conduct

See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
