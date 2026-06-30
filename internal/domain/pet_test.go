package domain

import (
	"testing"
)

func TestNewPet(t *testing.T) {
	p := NewPet("TestPet")
	if p.Name != "TestPet" {
		t.Errorf("expected name TestPet, got %s", p.Name)
	}
	if p.Stage != StageEgg {
		t.Errorf("expected stage egg, got %s", p.Stage)
	}
	if p.Energy != 50 {
		t.Errorf("expected energy 50, got %d", p.Energy)
	}
}

func TestApplyGoodEvent(t *testing.T) {
	p := NewPet("Test")
	initialEnergy := p.Energy
	p.ApplyGoodEvent(10, 5, 20)
	if p.Energy != clamp(initialEnergy+10, 0, 100) {
		t.Errorf("energy not updated correctly")
	}
	if p.Experience != 20 {
		t.Errorf("experience not updated, got %d", p.Experience)
	}
}

func TestApplyBadEvent(t *testing.T) {
	p := NewPet("Test")
	initialStress := p.Stress
	p.ApplyBadEvent(10, 5, 3)
	if p.Stress != clamp(initialStress+10, 0, 100) {
		t.Errorf("stress not updated correctly")
	}
}

func TestRecalculateMood(t *testing.T) {
	p := NewPet("Test")
	p.Energy = 100
	p.Stress = 0
	p.Hunger = 0
	p.RecalculateMood()
	if p.Mood != MoodEcstatic {
		t.Errorf("expected ecstatic, got %s", p.Mood)
	}

	p.Energy = 0
	p.Stress = 100
	p.Hunger = 100
	p.RecalculateMood()
	if p.Mood != MoodChaotic {
		t.Errorf("expected chaotic, got %s", p.Mood)
	}
}

func TestFeed(t *testing.T) {
	p := NewPet("Test")
	p.Hunger = 50
	p.Feed()
	if p.Hunger != 30 {
		t.Errorf("expected hunger 30, got %d", p.Hunger)
	}
}

func TestPlay(t *testing.T) {
	p := NewPet("Test")
	p.Stress = 30
	p.Play()
	if p.Stress != 15 {
		t.Errorf("expected stress 15, got %d", p.Stress)
	}
}

func TestClamp(t *testing.T) {
	if clamp(150, 0, 100) != 100 {
		t.Error("clamp max failed")
	}
	if clamp(-10, 0, 100) != 0 {
		t.Error("clamp min failed")
	}
	if clamp(50, 0, 100) != 50 {
		t.Error("clamp mid failed")
	}
}
