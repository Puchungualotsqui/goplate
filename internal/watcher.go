package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func RunWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	commands := [][]string{
		{"go", "run", "main.go"},
		{"templ", "generate"},
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
					continue
				}

				fmt.Println("Change detected:", event.Name)
				runCommands(commands)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher errors:", err)
		}
	}
}

func runCommands(cmds [][]string) {
	for _, cmdArgs := range cmds {
		fmt.Println("Running:", cmdArgs)
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error", err)
		}
	}
}
