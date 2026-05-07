package task

import "testing"

func TestAddTask(t *testing.T){
	tasks := []Task{}

	tasks = AddTask(tasks, "Learn Go", "testing")

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	task := tasks[0]

	if task.ID != 1 {
		t.Errorf("expected ID 1. got %d", task.ID)
	}

	if task.Name != "Learn Go" {
		t.Error("wrong task name")
	}

	if task.Description != "testing" {
		t.Error("wrong description")
	}

	if task.Done != false {
		t.Error("expected task not to be done")
	}
}

func TestDeleteTask(t *testing.T){
	tasks := []Task{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
		{ID: 3, Name: "Task 3"},
	}

	tasks, deleted := DeleteTask(tasks, 2)

	if !deleted {
		t.Fatal("expected task to be deleted")
	}

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].ID != 1 {
		t.Error("expeced first task ID to be 1")
	}

	if tasks[1].ID != 2 {
		t.Error("expeced reordered second task ID to be 2")
	}

	if tasks[1].Name != "Task 3" {
		t.Errorf("expected Task 3 to move into second position")
	}
}