package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCMD = &cobra.Command{
	Use:   "new [project name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Println("Creating project:", projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCMD)
}
