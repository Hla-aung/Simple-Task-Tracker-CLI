package task

import (
	"encoding/json"
	"os"
)

func SaveTasks(tasks []Task, filePath string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)

	return err
}

func LoadTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return []Task{}, nil
	}

	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}