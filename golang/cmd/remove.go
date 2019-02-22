package cmd

import (
	"fmt"
	"os"

	task "github.com/gdgtoledo/nodeGoCLI/golang/model"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(
		&description, "description", "d", defaultDescription, "Description of the task to be removed")
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Long:  `Remove a task, by its description`,
	Run: func(cmd *cobra.Command, args []string) {
		err := task.Remove(description)
		if err != nil {
			fmt.Println(color.RedString("Error removing task"))
			os.Exit(1)
		}

		err = task.List()
		if err != nil {
			os.Exit(1)
		}
	},
}
