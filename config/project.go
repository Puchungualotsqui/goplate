package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Puchungualotsqui/goplate/utils"

	"github.com/BurntSushi/toml"
)

type GoplateConfig struct {
	Project  string `toml:"project"`
	Module   string `toml:"module"`
	Tailwind bool   `toml:"tailwind"`
	DaisyUI  bool   `toml:"daisyui"`
}

func SaveConfig(projectPath string, cfg GoplateConfig) error {
	path, err := utils.ResolvePath(projectPath)
	if err != nil {
		return err
	}

	filePath := filepath.Join(path, "goplate.toml")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("failed to write goplate.toml: %w", err)
	}

	return nil
}

func LoadConfig(projectPath string) (GoplateConfig, error) {
	path, err := utils.ResolvePath(projectPath)
	if err != nil {
		return GoplateConfig{}, err
	}

	filePath := filepath.Join(path, "goplate.toml")
	var cfg GoplateConfig

	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to read goplate.toml: %w", err)
	}
	return cfg, nil
}
