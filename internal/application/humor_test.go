package application

import "testing"

func TestGetResponse(t *testing.T) {
	h := NewHumorEngine()

	categories := []string{"commit", "test", "lint", "dependency", "security", "idle", "push", "merge"}
	for _, cat := range categories {
		resp := h.GetResponse(cat)
		if resp == "" {
			t.Errorf("expected non-empty response for category %s", cat)
		}
	}
}

func TestGetResponseUnknownCategory(t *testing.T) {
	h := NewHumorEngine()
	resp := h.GetResponse("nonexistent")
	if resp == "" {
		t.Error("expected fallback message for unknown category")
	}
}

func TestGetFakeChaosMessage(t *testing.T) {
	msg := GetFakeChaosMessage()
	if msg == "" {
		t.Error("expected non-empty chaos message")
	}
}
