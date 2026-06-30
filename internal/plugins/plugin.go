package plugins

import (
	"context"

	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

// Plugin is the interface all Git-Toxagotchi plugins must implement.
// Plugins are observers — they cannot modify pet state directly.
type Plugin interface {
	Name() string
	OnStartup(ctx context.Context, pet *domain.Pet) error
	OnEvent(ctx context.Context, event *domain.Event, pet *domain.Pet) error
	OnEvolution(ctx context.Context, oldStage, newStage domain.Stage) error
	OnShutdown(ctx context.Context) error
}
