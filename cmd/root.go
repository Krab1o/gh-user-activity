package cmd

import (
	"fmt"
	"gh-user-activity/internal/activity"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use:   "github-activity",
	Short: "github-activity app",
	Run: func(cmd *cobra.Command, args []string) {
		activity.Activity(args[0])
	},
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}