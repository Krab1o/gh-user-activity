package cmd

import (
	"fmt"
	"gh-user-activity/internal/activity"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use:   "github-activity",
	Short: "github-activity app",
	Run: func(cmd *cobra.Command, args []string) {
		val, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		eventNumber := uint32(val)
		if (eventNumber > 100) {
			log.Fatal("The maximum number of requested events is 100!")
		} else {
			activity.Activity(args[0], eventNumber)
		}
		
	},
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}