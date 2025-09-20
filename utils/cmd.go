package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommands(cmds [][]string, dir string) error {
	for _, cmdArgs := range cmds {
		if err := RunCommand(cmdArgs, dir); err != nil {
			return err
		}
	}

	return nil
}

func RunCommand(cmdInput []string, dir string) error {
	fmt.Println("Running:", cmdInput)
	cmd := exec.Command(cmdInput[0], cmdInput[1:]...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
