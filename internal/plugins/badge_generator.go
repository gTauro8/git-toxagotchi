package plugins

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

// BadgeGeneratorPlugin generates a local SVG badge when the pet evolves.
// Future: post badge to GitHub Gist (opt-in, requires explicit user consent).
type BadgeGeneratorPlugin struct {
	OutputDir string
}

func (p *BadgeGeneratorPlugin) Name() string { return "badge-generator" }

func (p *BadgeGeneratorPlugin) OnStartup(_ context.Context, _ *domain.Pet) error { return nil }

func (p *BadgeGeneratorPlugin) OnEvent(_ context.Context, _ *domain.Event, _ *domain.Pet) error {
	return nil
}

func (p *BadgeGeneratorPlugin) OnEvolution(_ context.Context, _, newStage domain.Stage) error {
	if p.OutputDir == "" {
		home, _ := os.UserHomeDir()
		p.OutputDir = filepath.Join(home, ".local", "share", "git-toxagotchi")
	}
	if err := os.MkdirAll(p.OutputDir, 0755); err != nil {
		return fmt.Errorf("badge-generator: mkdir: %w", err)
	}
	path := filepath.Join(p.OutputDir, "badge.svg")
	svg := generateBadgeSVG(string(newStage))
	if err := os.WriteFile(path, []byte(svg), 0644); err != nil {
		return fmt.Errorf("badge-generator: write badge: %w", err)
	}
	fmt.Printf("📛 Badge updated: %s\n", path)
	return nil
}

func (p *BadgeGeneratorPlugin) OnShutdown(_ context.Context) error { return nil }

func generateBadgeSVG(stage string) string {
	label := "git-toxagotchi"
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="20">
  <linearGradient id="s" x2="0" y2="100%%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>
  <rect rx="3" width="200" height="20" fill="#555"/>
  <rect rx="3" x="120" width="80" height="20" fill="#4c1"/>
  <rect rx="3" width="200" height="20" fill="url(#s)"/>
  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text x="60" y="15" fill="#010101" fill-opacity=".3">%s</text>
    <text x="60" y="14">%s</text>
    <text x="160" y="15" fill="#010101" fill-opacity=".3">%s</text>
    <text x="160" y="14">%s</text>
  </g>
</svg>`, label, label, stage, stage)
}
