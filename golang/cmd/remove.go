package cmd

import (
	"log"

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
			log.Fatalln(color.RedString("Error removing task"))
		}

		tasks, _ := task.List()
		for _, t := range tasks {
			log.Println(color.YellowString(t.ToString()))
		}
	},
}
