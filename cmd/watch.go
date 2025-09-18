package cmd

import (
	"goplate/internal"

	"github.com/spf13/cobra"
)

var watchCMD = &cobra.Command{
	Use:   "watch",
	Short: "Watch files and run update commands on changes",
	Run: func(cmd *cobra.Command, args []string) {
		internal.RunWatcher()
	},
}

func init() {
	rootCmd.AddCommand(watchCMD)
}
