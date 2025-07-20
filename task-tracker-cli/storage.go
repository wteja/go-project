package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tasks: %w", err)
	}

	return tasks, nil
}

func saveTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling task: %w", err)
	}

	err = os.WriteFile(tasksFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing task to file: %w", err)
	}

	return nil
}
