package cmd

import (
	"log"

	task "github.com/gdgtoledo/nodeGoCLI/golang/model"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const defaultDescription = "Darle a Jhon Snow el trono de hierro"

func init() {
	createCmd.Flags().BoolVarP(
		&completed, "completed", "c", false, "Marks a task as completed (default false)")
	createCmd.Flags().StringVarP(
		&description, "description", "d", defaultDescription, "Sets description for a task")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a single task",
	Long:  `Create a single task. By default the task is created with false complete state`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := task.Create(description, completed)
		if err != nil {
			log.Fatalln(color.RedString("Error creating the task"))
		}
	},
}
