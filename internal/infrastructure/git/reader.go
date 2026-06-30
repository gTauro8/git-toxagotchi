package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) IsInsideGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	err := cmd.Run()
	return err == nil
}

func (r *Reader) GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (r *Reader) GetLastCommitMessage() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (r *Reader) GetBranchName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

type RepoStats struct {
	Branch      string
	CommitCount int
	IsClean     bool
}

func (r *Reader) GetRepoStats() (*RepoStats, error) {
	branch, err := r.GetBranchName()
	if err != nil {
		branch = "unknown"
	}

	cmd := exec.Command("git", "rev-list", "--count", "HEAD")
	out, _ := cmd.Output()
	count := 0
	fmt.Sscanf(strings.TrimSpace(string(out)), "%d", &count)

	statusCmd := exec.Command("git", "status", "--porcelain")
	statusOut, _ := statusCmd.Output()
	isClean := strings.TrimSpace(string(statusOut)) == ""

	return &RepoStats{
		Branch:      branch,
		CommitCount: count,
		IsClean:     isClean,
	}, nil
}
