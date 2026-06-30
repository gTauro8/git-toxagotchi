package application

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/sugar_petauro/git-toxagotchi/internal/domain"
)

var (
	secretPatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)(aws_access_key_id|aws_secret_access_key)\s*=\s*\S+`),
		regexp.MustCompile(`(?i)(password|passwd|pwd)\s*=\s*\S+`),
		regexp.MustCompile(`(?i)(token|secret|api_key|apikey)\s*=\s*\S+`),
		regexp.MustCompile(`-----BEGIN (RSA|EC|DSA|OPENSSH) PRIVATE KEY-----`),
		regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
	}

	conventionalCommitPattern = regexp.MustCompile(`^(feat|fix|docs|style|refactor|test|chore|perf|ci|build|revert)(\(.+\))?!?:\s.+`)
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) AnalyzeCommit(message, diff string) *domain.CommitAnalysis {
	analysis := &domain.CommitAnalysis{
		Message: message,
	}

	analysis.MessageQuality = a.scoreMessage(message)
	analysis.SecretsDetected = a.detectSecrets(diff)
	analysis.TodosAdded = a.countTodos(diff)
	analysis.DepsAdded = a.countDependencies(diff)
	analysis.DiffSize = a.countDiffLines(diff)
	analysis.FilesChanged = a.countFilesChanged(diff)

	return analysis
}

func (a *Analyzer) scoreMessage(msg string) int {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		return 0
	}

	score := 0

	// Length check
	if len(msg) >= 10 {
		score += 2
	}
	if len(msg) >= 20 {
		score += 1
	}

	// Conventional commit format
	if conventionalCommitPattern.MatchString(msg) {
		score += 4
	}

	// Imperative mood hint (starts with capital verb-like word)
	words := strings.Fields(msg)
	if len(words) > 0 {
		first := words[0]
		if len(first) > 0 && unicode.IsUpper(rune(first[0])) {
			score += 1
		}
	}

	// Not just "fix" or "wip"
	lower := strings.ToLower(strings.TrimSpace(msg))
	badMessages := []string{"fix", "wip", "update", "changes", "misc", "temp", "test"}
	isBad := false
	for _, bad := range badMessages {
		if lower == bad {
			isBad = true
			break
		}
	}
	if !isBad {
		score += 2
	}

	if score > 10 {
		score = 10
	}
	return score
}

func (a *Analyzer) detectSecrets(diff string) bool {
	for _, pattern := range secretPatterns {
		if pattern.MatchString(diff) {
			return true
		}
	}
	return false
}

func (a *Analyzer) countTodos(diff string) int {
	count := 0
	todoPattern := regexp.MustCompile(`(?i)^\+.*(TODO|FIXME|HACK|XXX)`)
	for _, line := range strings.Split(diff, "\n") {
		if todoPattern.MatchString(line) {
			count++
		}
	}
	return count
}

func (a *Analyzer) countDependencies(diff string) int {
	count := 0
	lines := strings.Split(diff, "\n")
	inDepsFile := false
	for _, line := range lines {
		if strings.HasPrefix(line, "+++") {
			inDepsFile = strings.Contains(line, "go.mod") ||
				strings.Contains(line, "package.json") ||
				strings.Contains(line, "requirements.txt") ||
				strings.Contains(line, "Gemfile") ||
				strings.Contains(line, "Cargo.toml")
		}
		if inDepsFile && strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			count++
		}
	}
	return count
}

func (a *Analyzer) countDiffLines(diff string) int {
	added, removed := 0, 0
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			added++
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			removed++
		}
	}
	return added + removed
}

func (a *Analyzer) countFilesChanged(diff string) int {
	count := 0
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "diff --git") {
			count++
		}
	}
	return count
}
