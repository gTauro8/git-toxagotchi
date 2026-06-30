package domain

import "time"

type AchievementID string

const (
	AchFirstCommit           AchievementID = "first_commit"
	AchTenGoodCommits        AchievementID = "ten_good_commits"
	AchHundredTestsPassed    AchievementID = "hundred_tests_passed"
	AchNoForcePushWeek       AchievementID = "no_force_push_week"
	AchSecretHunter          AchievementID = "secret_hunter"
	AchDocumentationEnjoyer  AchievementID = "documentation_enjoyer"
	AchDependencyDiet        AchievementID = "dependency_diet"
	AchMergeConflictSurvivor AchievementID = "merge_conflict_survivor"
)

type Achievement struct {
	ID          AchievementID
	Name        string
	Description string
	Unlocked    bool
	UnlockedAt  *time.Time
	Icon        string
}

func AllAchievements() []Achievement {
	return []Achievement{
		{ID: AchFirstCommit, Name: "First Commit", Description: "Made your very first commit", Icon: "🐣"},
		{ID: AchTenGoodCommits, Name: "10 Good Commits", Description: "10 commits with quality messages", Icon: "✨"},
		{ID: AchHundredTestsPassed, Name: "100 Tests Passed", Description: "100 tests passed", Icon: "🧪"},
		{ID: AchNoForcePushWeek, Name: "No Force Push Week", Description: "Survived a week without force pushing", Icon: "🕊️"},
		{ID: AchSecretHunter, Name: "Secret Hunter", Description: "Detected and avoided committing a secret", Icon: "🔍"},
		{ID: AchDocumentationEnjoyer, Name: "Documentation Enjoyer", Description: "Updated docs 5 times", Icon: "📖"},
		{ID: AchDependencyDiet, Name: "Dependency Diet", Description: "No new dependencies for a week", Icon: "🥗"},
		{ID: AchMergeConflictSurvivor, Name: "Merge Conflict Survivor", Description: "Survived a merge conflict", Icon: "⚔️"},
	}
}
