package domain

import "testing"

func TestShouldEvolve(t *testing.T) {
	p := NewPet("Test")
	p.Experience = 0
	_, ok := ShouldEvolve(p)
	if ok {
		t.Error("should not evolve with 0 exp")
	}

	p.Experience = 100
	next, ok := ShouldEvolve(p)
	if !ok {
		t.Error("should evolve with 100 exp")
	}
	if next != StageBlob {
		t.Errorf("expected blob, got %s", next)
	}
}

func TestNextStage(t *testing.T) {
	next, ok := NextStage(StageEgg)
	if !ok || next != StageBlob {
		t.Errorf("expected blob, got %s, ok=%v", next, ok)
	}

	_, ok = NextStage(StageLegendaryMaintainer)
	if ok {
		t.Error("legendary maintainer has no next stage")
	}
}
