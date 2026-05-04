package task

import "fmt"

type Task struct {
	ID   int
	Name string
	Description string
	Done bool
}

func AddTask(tasks []Task, name string, description string) []Task {
	id := len(tasks) + 1
	return append(tasks, Task{
		ID:   id,
		Name: name,
		Description: description,
		Done: false,
	})
}

func ReorderTasks(tasks []Task) []Task {
	for i := range tasks {
		tasks[i].ID = i + 1
	}

	return tasks
}

func DeleteTask(tasks []Task, id int) ([]Task, bool) {
	for i, task := range tasks {
		if id == task.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			tasks = ReorderTasks(tasks)
			return tasks, true
		}
	}
	return tasks, false
}

func ListTasks(tasks []Task) {
	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %v %v\n", task.ID, task.Name, status)
		if(task.Description != ""){
			fmt.Println("  ",task.Description)
		}
	}
}

func MarkDone(tasks []Task, id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return true
		}
	}

	return false
}