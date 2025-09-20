package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Puchungualotsqui/goplate/config"
	"github.com/Puchungualotsqui/goplate/internal"
	"github.com/Puchungualotsqui/goplate/internal/checks"
	"github.com/Puchungualotsqui/goplate/internal/skeleton"
	"github.com/Puchungualotsqui/goplate/utils"

	"github.com/spf13/cobra"
)

var newCMD = &cobra.Command{
	Use:   "new [project name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GoplateConfig{}

		projectName := args[0]

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("❌ Could not get working directory:", err)
			return
		}

		projectPath := filepath.Join(cwd, projectName)

		if err := checks.EnsureTemplInstalled(); err != nil {
			fmt.Println(err)
			return
		}

		if err := internal.CreateSkeleton(projectName, projectName, skeleton.DefaultSkeleton); err != nil {
			fmt.Println("Error:", err)
			return
		}

		if err := checks.EnsureTailwindInstalled(projectPath, projectName, &cfg); err != nil {
			fmt.Println("Error:", err)
			return
		}

		if err := utils.RunCommands([][]string{
			{"go", "mod", "init", projectName},
			{"go", "mod", "tidy"},
			{"go", "get", "github.com/a-h/templ"},
			{"templ", "generate"},
		}, projectPath); err != nil {
			fmt.Println("Error initializing go project:", err)
		}

		fmt.Println("Project created at", projectName)

		if err := config.SaveConfig(projectPath, cfg); err != nil {
			fmt.Println("Error saving config:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCMD)
}
