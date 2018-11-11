package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your tasks",
	Long:  `List all your tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello List!")
	},
}
