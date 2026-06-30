package config

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.PetName != "Toxi" {
		t.Errorf("expected PetName Toxi, got %s", cfg.PetName)
	}
	if cfg.Theme != "default" {
		t.Errorf("expected theme default, got %s", cfg.Theme)
	}
	if cfg.LLMEnabled {
		t.Error("expected LLMEnabled false by default")
	}
	if cfg.HookBlocking {
		t.Error("expected HookBlocking false by default")
	}
	if cfg.DBPath == "" {
		t.Error("expected non-empty DBPath")
	}
}

func TestConfigPath(t *testing.T) {
	path := Path()
	if path == "" {
		t.Error("config path should not be empty")
	}
}
