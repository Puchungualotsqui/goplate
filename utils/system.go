package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func DetectPlatform() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}

func ResolvePlatform(goos, arch string) (string, error) {
	switch goos {
	case "linux":
		switch arch {
		case "amd64":
			return "linux-x64", nil
		case "arm64":
			return "linux-arm64", nil
		}
	case "darwin":
		switch arch {
		case "amd64":
			return "macos-x64", nil
		case "arm64":
			return "macos-arm64", nil
		}
	case "windows":
		return "windows-x64.exe", nil
	}
	return "", fmt.Errorf("unsupported platform: %s/%s", goos, arch)
}

func DownloadFile(url, target string, makeExec bool) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", url, err)
	}
	defer resp.Body.Close()

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	if makeExec {
		if err := os.Chmod(target, 0755); err != nil {
			return err
		}
	}

	return nil
}

func ResolvePath(projectPath string) (string, error) {
	if projectPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get working dir: %w", err)
		}
		return cwd, nil
	}
	return projectPath, nil
}
