# Plugin Guide

Git-Toxagotchi supports plugins via a simple Go interface.

## Plugin Interface

```go
type Plugin interface {
    Name() string
    OnStartup(ctx context.Context, pet *domain.Pet) error
    OnEvent(ctx context.Context, event *domain.Event, pet *domain.Pet) error
    OnEvolution(ctx context.Context, oldStage, newStage domain.Stage) error
    OnShutdown(ctx context.Context) error
}
```

## Built-in Plugins

### ConsoleLoggerPlugin
Logs all events to stderr. Useful for debugging.

### FakeChaosPlugin
Randomly prints fake-chaotic messages after bad events. All messages are harmless text.

### BadgeGeneratorPlugin (stub)
Generates a local SVG badge showing pet status. Future: post to GitHub Gist.

## Writing a Plugin

```go
package myplugin

import (
    "context"
    "fmt"
    "github.com/sugar_petauro/git-toxagotchi/internal/domain"
    "github.com/sugar_petauro/git-toxagotchi/internal/plugins"
)

type MyPlugin struct{}

func (p *MyPlugin) Name() string { return "my-plugin" }

func (p *MyPlugin) OnStartup(ctx context.Context, pet *domain.Pet) error {
    fmt.Println("Plugin started!")
    return nil
}

func (p *MyPlugin) OnEvent(ctx context.Context, event *domain.Event, pet *domain.Pet) error {
    fmt.Printf("Event: %s\n", event.Type)
    return nil
}

func (p *MyPlugin) OnEvolution(ctx context.Context, old, new domain.Stage) error {
    fmt.Printf("Evolved: %s -> %s\n", old, new)
    return nil
}

func (p *MyPlugin) OnShutdown(ctx context.Context) error { return nil }
```

## Plugin Constraints

- Plugins are **observers** — they cannot modify pet state directly
- Plugins must not perform destructive operations
- Plugins must not make network calls without user consent
- Plugin errors are logged but do not crash the application

## Registering Plugins

Register in the application service:

```go
svc.RegisterPlugin(&myplugin.MyPlugin{})
```
