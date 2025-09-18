package cmd

import (
	"fmt"

	"goplate/internal"
	"goplate/internal/skeleton"

	"github.com/spf13/cobra"
)

var newCMD = &cobra.Command{
	Use:   "new [project name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if err := internal.CreateProject(projectName, skeleton.DefaultSkeleton); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Project created at", projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCMD)
}
