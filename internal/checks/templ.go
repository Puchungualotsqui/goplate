package checks

import (
	"fmt"
	"os/exec"

	"github.com/Puchungualotsqui/goplate/utils"
)

func EnsureTemplInstalled() error {
	_, err := exec.LookPath("templ")
	if err == nil {
		fmt.Println("templ is installed")
		return nil
	}

	fmt.Println("templ was not found in Path.")
	fmt.Print("Do you want to install it now? [y/N]: ")

	var input string
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {
		fmt.Println("Installing templ...")
		if err := utils.RunCommand([]string{"go", "install", "github.com/a-h/templ/cmd/templ@latest"}, ""); err != nil {
			return fmt.Errorf("Failed installing templ: %w", err)
		}
		fmt.Println("templ installed successfully! (make sure your $GOPATH/bin is in PATH)")
		return nil
	}

	return fmt.Errorf("templ not installed, please install manually: go install github.com/a-h/templ/cmd/templ@latest")
}
