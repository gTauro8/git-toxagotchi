package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PetName      string `yaml:"pet_name"`
	Theme        string `yaml:"theme"`
	LLMEnabled   bool   `yaml:"llm_enabled"`
	HookBlocking bool   `yaml:"hook_blocking"`
	DBPath       string `yaml:"db_path"`
}

func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		PetName:      "Toxi",
		Theme:        "default",
		LLMEnabled:   false,
		HookBlocking: false,
		DBPath:       filepath.Join(home, ".config", "git-toxagotchi", "pet.db"),
	}
}

func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "git-toxagotchi", "config.yaml")
}

func Load() (*Config, error) {
	cfg := DefaultConfig()
	path := ConfigPath()

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func Save(cfg *Config) error {
	path := ConfigPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
