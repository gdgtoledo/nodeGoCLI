package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var completed bool
var description string

var rootCmd = &cobra.Command{
	Use:   "gdg-cli",
	Short: "gdg-cli (Google Developer Group CLI) It's a simple and functional Golang To-Do List.",
	Long: `gdg-cli (Google Developer Group CLI) It's a simple and functional Golang To-Do List.
				built with love by GDGToledo and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.RedString("", err))
		os.Exit(1)
	}
}
