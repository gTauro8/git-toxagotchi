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

	cases := []struct {
		name string
		diff string
		want bool
	}{
		{"AWS key ID", `+AKIAIOSFODNN7EXAMPLE`, true},
		{"AWS env var", `+AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE`, true},
		{"RSA private key", `+-----BEGIN RSA PRIVATE KEY-----`, true},
		{"OPENSSH private key", `+-----BEGIN OPENSSH PRIVATE KEY-----`, true},
		{"GitHub token ghp_", "+token = ghp" + "_abcdefghijklmnopqrstuvwxyz123456789012", true},
		{"Stripe live key", "+STRIPE_KEY=sk" + "_live_abcdefghijklmnopqrstuvwx", true},
		{"Google API key", "+key = AIza" + "SyD-abcdefghijklmnopqrstuvwxyz1234567", true},
		{"Slack token", "+token = xoxb" + "-123456789012-123456789012-abcdef", true},
		{"SendGrid key", "+SG." + "aaaaaaaaaaaaaaaaaaaaaa.bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", true},
		{"npm token", "+authToken=npm" + "_abcdefghijklmnopqrstuvwxyz123456789012", true},
		{"generic password", `+password = "mysupersecret123"`, true},
		{".env file staged", "+++ b/config/.env", true},
		{"clean Go code", `+fmt.Println("hello world")`, false},
		{"clean comment", `+// this is a normal comment`, false},
		{"test token in test file", `+testToken := "test_value"`, false},
	}

	for _, tc := range cases {
		got := a.detectSecrets(tc.diff)
		if got != tc.want {
			t.Errorf("[%s] detectSecrets(%q) = %v, want %v", tc.name, tc.diff, got, tc.want)
		}
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
