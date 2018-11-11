package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update single task",
	Long:  `Update the complete state for a task`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello Update!")
	},
}
