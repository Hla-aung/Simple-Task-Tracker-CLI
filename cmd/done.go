package cmd

import (
	"fmt"
	"strconv"

	"example.com/struct-list/task"
	"github.com/spf13/cobra"
)

var doneCommand = &cobra.Command{
	Use:   "done",
	Short: "Mark task done",
	Run:   runDoneCommand,
}

func init() {
	rootCommand.AddCommand(doneCommand)
}

func runDoneCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go run . done \"Task ID\"")
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
	found := task.MarkDone(loadedTasks, id)
	if !found {
		fmt.Println("Task not found with given ID: ", id)
		return
	}
	err = task.SaveTasks(loadedTasks, filePath)
	if err != nil {
		fmt.Println("Error saving tasks: ", err)
		return
	}
	fmt.Println("Task done")
}
