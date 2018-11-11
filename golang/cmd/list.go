package cmd

import (
	"log"

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
		tasks, err := task.List()
		if err != nil {
			log.Fatalln("Error listing tasks")
		}

		for _, t := range tasks {
			log.Println(t)
		}
	},
}
