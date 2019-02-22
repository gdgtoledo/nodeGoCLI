package cmd

import (
	"os"

	task "github.com/gdgtoledo/nodeGoCLI/golang/model"

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
		err := task.List()
		if err != nil {
			os.Exit(1)
		}
	},
}
