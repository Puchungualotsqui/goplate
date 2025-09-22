package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/Puchungualotsqui/goplate/config"
	"github.com/Puchungualotsqui/goplate/utils"

	"github.com/bmatcuk/doublestar/v4"
)

var mu sync.Mutex
var timer *time.Timer

func scheduleRun(buildCommands [][]string) {
	mu.Lock()
	defer mu.Unlock()
	if timer != nil {
		timer.Stop()
	}
	timer = time.AfterFunc(300*time.Millisecond, func() {
		// Run build steps
		utils.RunCommands(buildCommands, "")
		// Restart server
		runServer()
	})
}

func RunWatcher() {
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	buildCommands := [][]string{
		{"templ", "generate"},
	}

	if cfg.Tailwind {
		buildCommands = append([][]string{
			{"./static/css/tailwindcss", "-c", "./static/css/tailwind.config.js", "-i", "./static/css/input.css", "-o", "./static/css/output.css"},
		}, buildCommands...)
	}

	watchDirs := []string{"."}
	for _, dir := range watchDirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() {
				watcher.Add(path)
			}
			return nil
		})
	}

	scheduleRun(buildCommands)

	fmt.Println("Watching for changes...")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
				// Ignore folder creation
				info, err := os.Stat(event.Name)
				if err == nil && info.IsDir() {
					if event.Op&fsnotify.Create != 0 {
						watcher.Add(event.Name)
						fmt.Println("📂 Added new dir to watch:", event.Name)
					}
					continue
				}

				if ShouldIgnore(event.Name, cfg.IgnorePatterns) {
					continue
				}

				fmt.Println("Change detected:", event.Name)
				scheduleRun(buildCommands)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher errors:", err)
		}
	}
}

func ShouldIgnore(path string, patterns []string) bool {
	// Normalize slashes for cross-platform matching
	path = filepath.ToSlash(filepath.Clean(path))

	for _, pattern := range patterns {
		// Also normalize patterns (so they work on Windows too)
		pattern = filepath.ToSlash(pattern)

		matched, err := doublestar.PathMatch(pattern, path)
		if err != nil {
			continue
		}
		if matched {
			return true
		}
	}
	return false
}
