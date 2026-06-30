package plugins

import (
	"context"
	"fmt"
	"time"

	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

// ConsoleLoggerPlugin logs all pet events to stderr. Useful for debugging.
type ConsoleLoggerPlugin struct{}

func (p *ConsoleLoggerPlugin) Name() string { return "console-logger" }

func (p *ConsoleLoggerPlugin) OnStartup(_ context.Context, pet *domain.Pet) error {
	fmt.Printf("[%s] console-logger: startup, pet=%s stage=%s\n", ts(), pet.Name, pet.Stage)
	return nil
}

func (p *ConsoleLoggerPlugin) OnEvent(_ context.Context, event *domain.Event, pet *domain.Pet) error {
	fmt.Printf("[%s] console-logger: event=%s pet=%s mood=%s\n", ts(), event.Type, pet.Name, pet.Mood)
	return nil
}

func (p *ConsoleLoggerPlugin) OnEvolution(_ context.Context, old, new domain.Stage) error {
	fmt.Printf("[%s] console-logger: evolution %s -> %s\n", ts(), old, new)
	return nil
}

func (p *ConsoleLoggerPlugin) OnShutdown(_ context.Context) error {
	fmt.Printf("[%s] console-logger: shutdown\n", ts())
	return nil
}

func ts() string {
	return time.Now().Format("15:04:05")
}
