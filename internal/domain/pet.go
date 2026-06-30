package domain

import (
	"fmt"
	"time"
)

type Stage string

const (
	StageEgg                 Stage = "egg"
	StageBlob                Stage = "blob"
	StageGremlin             Stage = "gremlin"
	StageDeveloperCat        Stage = "developer_cat"
	StageTerminalWizard      Stage = "terminal_wizard"
	StageKernelPenguin       Stage = "kernel_penguin"
	StageGitDragon           Stage = "git_dragon"
	StageLegendaryMaintainer Stage = "legendary_maintainer"
)

type Mood string

const (
	MoodEcstatic  Mood = "ecstatic"
	MoodHappy     Mood = "happy"
	MoodContent   Mood = "content"
	MoodNeutral   Mood = "neutral"
	MoodAnnoyed   Mood = "annoyed"
	MoodGrumpy    Mood = "grumpy"
	MoodMiserable Mood = "miserable"
	MoodChaotic   Mood = "chaotic"
)

type Pet struct {
	ID                string
	Name              string
	Species           string
	Stage             Stage
	Energy            int
	Hunger            int
	Stress            int
	Trust             int
	Chaos             int
	Experience        int
	Age               int
	LastInteractionAt time.Time
	Mood              Mood
}

func NewPet(name string) *Pet {
	return &Pet{
		ID:                generateID(),
		Name:              name,
		Species:           "Toxagotchi",
		Stage:             StageEgg,
		Energy:            50,
		Hunger:            30,
		Stress:            10,
		Trust:             50,
		Chaos:             0,
		Experience:        0,
		Age:               0,
		LastInteractionAt: time.Now(),
		Mood:              MoodContent,
	}
}

func (p *Pet) RecalculateMood() {
	score := p.Energy - p.Stress - p.Hunger/2
	switch {
	case score >= 80:
		p.Mood = MoodEcstatic
	case score >= 60:
		p.Mood = MoodHappy
	case score >= 40:
		p.Mood = MoodContent
	case score >= 20:
		p.Mood = MoodNeutral
	case score >= 0:
		p.Mood = MoodAnnoyed
	case score >= -20:
		p.Mood = MoodGrumpy
	case score >= -40:
		p.Mood = MoodMiserable
	default:
		p.Mood = MoodChaotic
	}
}

func (p *Pet) ApplyGoodEvent(energyDelta, trustDelta, expDelta int) {
	p.Energy = clamp(p.Energy+energyDelta, 0, 100)
	p.Trust = clamp(p.Trust+trustDelta, 0, 100)
	p.Experience += expDelta
	p.Stress = clamp(p.Stress-5, 0, 100)
	p.RecalculateMood()
}

func (p *Pet) ApplyBadEvent(stressDelta, hungerDelta, chaosDelta int) {
	p.Stress = clamp(p.Stress+stressDelta, 0, 100)
	p.Hunger = clamp(p.Hunger+hungerDelta, 0, 100)
	p.Chaos = clamp(p.Chaos+chaosDelta, 0, 100)
	p.Energy = clamp(p.Energy-5, 0, 100)
	p.RecalculateMood()
}

func (p *Pet) Feed() {
	p.Hunger = clamp(p.Hunger-20, 0, 100)
	p.Energy = clamp(p.Energy+10, 0, 100)
	p.RecalculateMood()
}

func (p *Pet) Play() {
	p.Energy = clamp(p.Energy-10, 0, 100)
	p.Stress = clamp(p.Stress-15, 0, 100)
	p.Trust = clamp(p.Trust+5, 0, 100)
	p.RecalculateMood()
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func generateID() string {
	return fmt.Sprintf("pet_%d", time.Now().UnixNano())
}
