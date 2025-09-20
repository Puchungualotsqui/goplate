package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileTemplate struct {
	Path    string
	Content string
	IsDir   bool
}

func CreateSkeleton(basePath, moduleName string, skeleton []FileTemplate) error {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return fmt.Errorf("failed to create project root %s: %w:", basePath, err)
	}
	for _, item := range skeleton {
		targetPath := filepath.Join(basePath, item.Path)

		if item.IsDir {
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("failed to create dir %s: %w", targetPath, err)
			}
			continue
		}

		content := strings.ReplaceAll(item.Content, "{{MODULE}}", moduleName)
		content = strings.TrimLeft(item.Content, "\n\r\t ")

		dir := filepath.Dir(targetPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := os.WriteFile(targetPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", targetPath, err)
		}
	}
	return nil
}
