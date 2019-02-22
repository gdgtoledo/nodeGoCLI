package cmd

import (
	"fmt"
	"os"

	task "github.com/gdgtoledo/nodeGoCLI/golang/model"

	"github.com/fatih/color"
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
			fmt.Println(color.RedString("Error listing tasks"))
			os.Exit(1)
		}

		for _, t := range tasks {
			fmt.Println(color.YellowString(t.ToString()))
		}
	},
}
