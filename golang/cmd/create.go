package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a single task",
	Long:  `Create a single task. By default the task is created with false complete state`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello Create!")
	},
}
