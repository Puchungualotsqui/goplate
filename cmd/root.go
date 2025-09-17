package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goplate",
	Short: "Goplate - scaffold Go+Templ+HTMX apps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Goplate!")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
