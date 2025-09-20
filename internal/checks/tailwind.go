package checks

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Puchungualotsqui/goplate/config"
	"github.com/Puchungualotsqui/goplate/internal"
	"github.com/Puchungualotsqui/goplate/internal/skeleton"
	"github.com/Puchungualotsqui/goplate/utils"
)

func EnsureTailwindInstalled(projectPath, basePath string, cfg *config.GoplateConfig) error {
	fmt.Print("Do you want to install tailwind? [y/N]: ")

	var input string
	fmt.Scanln(&input)

	if input != "y" && input != "Y" {
		return nil
	}

	if err := internal.CreateSkeleton(basePath, "", skeleton.TailwindSkeleton); err != nil {
		return fmt.Errorf("Error creating tailwind folders: %w", err)
	}

	platform, err := utils.ResolvePlatform(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-%s", platform)
	targetFile := filepath.Join(projectPath, "static", "css", "tailwindcss")
	if runtime.GOOS == "windows" {
		targetFile += ".exe"
	}

	if err := utils.DownloadFile(url, targetFile, runtime.GOOS != "windows"); err != nil {
		return err
	}

	cfg.Tailwind = true
	fmt.Println("TailwindCSS installed locally at", targetFile)
	return nil
}
