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

	// Kill old server if running
	if serverCmd != nil && serverCmd.Process != nil {
		fmt.Println("🛑 Stopping old server...")
		_ = serverCmd.Process.Kill()
		serverCmd.Wait() // wait to release the port
	}

	// Start new server
	fmt.Println("🚀 Starting server...")
	serverCmd = exec.Command("go", "run", ".")
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr
	if err := serverCmd.Start(); err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
