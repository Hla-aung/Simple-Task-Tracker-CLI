package cmd

import (
	"fmt"

	"example.com/struct-list/task"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Run:   runAddCommand,
}

var description string

func init() {
	addCommand.Flags().StringVarP(&description, "description", "d", "", "add description to task")
	rootCommand.AddCommand(addCommand)
}

func runAddCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go run . add \"task name\" [\"description\"]")
		return
	}
	loadedTasks, err := task.LoadTasks(filePath)
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
	}
	tasks := task.AddTask(loadedTasks, args[0], description)
	err = task.SaveTasks(tasks, filePath)
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		return
	}
	fmt.Println("Task added")
}
