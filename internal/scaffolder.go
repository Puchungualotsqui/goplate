package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileTemplate struct {
	Path    string
	Content string
	IsDir   bool
}

func CreateProject(basePath string, skeleton []FileTemplate) error {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return fmt.Errorf("failed to create project root %s: %w:", basePath, err)
	}
	for _, item := range skeleton {
		targetPath := filepath.Join(basePath, item.Path)

		if item.IsDir {
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("failed to create dir %s: %w", targetPath, err)
			}
		} else {
			dir := filepath.Dir(targetPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if err := os.WriteFile(targetPath, []byte(item.Content), 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
		}
	}
	return nil
}
