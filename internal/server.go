package internal

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var serverCmd *exec.Cmd
var serverMu sync.Mutex

func runServer() {
	serverMu.Lock()
	defer serverMu.Unlock()

	// Kill old server
	if serverCmd != nil && serverCmd.Process != nil {
		fmt.Println("🛑 Stopping old server...")
		_ = serverCmd.Process.Kill()
		serverCmd.Wait()
	}

	// Rebuild
	build := exec.Command("go", "build", "-o", "./bin/server", ".")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Println("❌ Build failed:", err)
		return
	}

	// Start new binary
	fmt.Println("🚀 Starting server...")
	serverCmd = exec.Command("./bin/server")
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr
	if err := serverCmd.Start(); err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
