package main

import (
	"github.com/spf13/cobra"
)

func main() {

	// chronoCmd represents the chrono command
	var chronoCmd = &cobra.Command{
		Use:   "chrono",
		Short: "Starts a stopwatch.",
		Long:  `Starts a stopwatch clock. Pressing the spacebar result take a lap time.`,
		Run: func(cmd *cobra.Command, args []string) {
			chrono()
		},
	}
	var rootCmd = &cobra.Command{Use: "timer"}
	rootCmd.AddCommand(chronoCmd)
	rootCmd.Execute()
}
