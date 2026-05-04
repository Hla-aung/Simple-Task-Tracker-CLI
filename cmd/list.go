package cmd

import (
	"fmt"

	"example.com/struct-list/task"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run:   runListCommand,
}

func init() {
	rootCommand.AddCommand(listCommand)
}

func runListCommand(cmd *cobra.Command, args []string) {
	tasks, err := task.LoadTasks(filePath)

	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}

	task.ListTasks(tasks)
}
