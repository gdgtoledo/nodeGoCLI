package cmd

import (
	"fmt"
	"os"

	task "github.com/gdgtoledo/nodeGoCLI/golang/model"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	updateCmd.Flags().BoolVarP(
		&completed, "completed", "c", false, "Marks a task as completed (default false)")
	updateCmd.Flags().StringVarP(
		&description, "description", "d", defaultDescription, "Sets description for a task")

	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update single task",
	Long:  `Update the complete state for a task`,
	Run: func(cmd *cobra.Command, args []string) {
		err := task.Update(description, completed)
		if err != nil {
			fmt.Println(color.RedString(err.Error()))
			os.Exit(1)
		}

		err = task.List()
		if err != nil {
			os.Exit(1)
		}
	},
}
