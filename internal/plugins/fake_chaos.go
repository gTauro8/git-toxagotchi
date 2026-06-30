package plugins

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gTauro8/git-toxagotchi/internal/domain"
)

var chaosMessages = []string{
	"Deleting repository... just kidding. Or am I? (I'm not.)",
	"Running `rm -rf /`... nah, I like your files.",
	"Pet tried to eat node_modules and got indigestion.",
	"Formatting disk... no, I'm not that kind of monster.",
	"Uploading your commits to the dark web... psych!",
	"Pet attempted to rewrite history. Git refused. Both are fine.",
	"Pushing directly to main... in my dreams.",
	"Pet ate your TODO list. Now it's a DONE list. You're welcome.",
	"Squashing all commits into one... I thought about it.",
	"Pet tried to git blame you. The mirror blocked it.",
	"Initiating DEFCON 1... actually just making tea.",
	"Running 'sudo rm -rf --no-preserve-root'... just checking if you panicked.",
	"Pet found your .env file. It is sworn to secrecy.",
	"Deploying to production directly from localhost... I could never.",
	"Deleting all branches except master... gotcha.",
}

// FakeChaosPlugin prints harmless fake-chaotic messages after bad events.
// It never performs any actual operations.
type FakeChaosPlugin struct{}

func (p *FakeChaosPlugin) Name() string { return "fake-chaos" }

func (p *FakeChaosPlugin) OnStartup(_ context.Context, _ *domain.Pet) error { return nil }

func (p *FakeChaosPlugin) OnEvent(_ context.Context, event *domain.Event, pet *domain.Pet) error {
	bad := event.Type == domain.EventForcePush ||
		event.Type == domain.EventSecretDetected ||
		event.Type == domain.EventTestFailed
	if bad && rand.Float32() < 0.4 {
		fmt.Printf("\n⚡ %s says: %s\n\n", pet.Name, chaosMessages[rand.Intn(len(chaosMessages))])
	}
	return nil
}

func (p *FakeChaosPlugin) OnEvolution(_ context.Context, _, _ domain.Stage) error { return nil }

func (p *FakeChaosPlugin) OnShutdown(_ context.Context) error { return nil }
