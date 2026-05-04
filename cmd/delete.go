package cmd

import (
	"fmt"
	"strconv"

	"example.com/struct-list/task"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run:   runDeleteCommand,
}

func init() {
	rootCommand.AddCommand(deleteCommand)
}

func runDeleteCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go run . delete \"Task ID\"")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error converting to number ID")
		return
	}
	loadedTasks, err := task.LoadTasks(filePath)
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
	}
	tasks, found := task.DeleteTask(loadedTasks, id)
	if !found {
		fmt.Println("Task not found with given ID: ", id)
		return
	}
	err = task.SaveTasks(tasks, filePath)
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		return
	}
	fmt.Println("Task deleted")
}
