package application

import "testing"

func TestScoreMessage(t *testing.T) {
	a := NewAnalyzer()

	tests := []struct {
		msg      string
		minScore int
	}{
		{"feat: add user authentication", 7},
		{"fix", 0},
		{"wip", 0},
		{"", 0},
		{"fix: correct null pointer dereference in handler", 7},
	}

	for _, tt := range tests {
		score := a.scoreMessage(tt.msg)
		if score < tt.minScore {
			t.Errorf("message %q: expected score >= %d, got %d", tt.msg, tt.minScore, score)
		}
	}
}

func TestDetectSecrets(t *testing.T) {
	a := NewAnalyzer()

	withSecret := `+AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE`
	if !a.detectSecrets(withSecret) {
		t.Error("should detect AWS key")
	}

	clean := `+fmt.Println("hello world")`
	if a.detectSecrets(clean) {
		t.Error("should not detect secret in clean diff")
	}

	privateKey := `+-----BEGIN RSA PRIVATE KEY-----`
	if !a.detectSecrets(privateKey) {
		t.Error("should detect private key")
	}
}

func TestCountTodos(t *testing.T) {
	a := NewAnalyzer()
	diff := `+// TODO: fix this later
+fmt.Println("hello")
+// FIXME: broken
-old code`
	count := a.countTodos(diff)
	if count != 2 {
		t.Errorf("expected 2 todos, got %d", count)
	}
}

func TestCountDiffLines(t *testing.T) {
	a := NewAnalyzer()
	diff := `+added line 1
+added line 2
-removed line
 context line`
	size := a.countDiffLines(diff)
	if size != 3 {
		t.Errorf("expected 3 diff lines, got %d", size)
	}
}
