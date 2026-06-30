package domain

var StageOrder = []Stage{
	StageEgg,
	StageBlob,
	StageGremlin,
	StageDeveloperCat,
	StageTerminalWizard,
	StageKernelPenguin,
	StageGitDragon,
	StageLegendaryMaintainer,
}

var StageExperienceThresholds = map[Stage]int{
	StageEgg:                 0,
	StageBlob:                100,
	StageGremlin:             300,
	StageDeveloperCat:        600,
	StageTerminalWizard:      1000,
	StageKernelPenguin:       1500,
	StageGitDragon:           2200,
	StageLegendaryMaintainer: 3000,
}

func NextStage(current Stage) (Stage, bool) {
	for i, s := range StageOrder {
		if s == current && i < len(StageOrder)-1 {
			return StageOrder[i+1], true
		}
	}
	return current, false
}

func ShouldEvolve(p *Pet) (Stage, bool) {
	next, ok := NextStage(p.Stage)
	if !ok {
		return p.Stage, false
	}
	if p.Experience >= StageExperienceThresholds[next] {
		return next, true
	}
	return p.Stage, false
}
