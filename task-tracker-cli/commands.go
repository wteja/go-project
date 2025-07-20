package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for i, task := range tasks {
		fmt.Printf("%d: %s [%s]\n", i+1, task.Description, task.Status)
	}
}

func addTask() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task-cli add <task_description>")
		os.Exit(1)
	}

	description := os.Args[2]
	if description == "" {
		fmt.Println("Task description cannot be empty.")
		os.Exit(1)
	}

	tasksList, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	nextID := getNextID(tasksList)

	task := Task{
		ID:          nextID,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	tasksList = append(tasksList, task)

	err = saveTask(tasksList)
	if err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", nextID)
}

func updateTask() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task-cli update <task_id> <new_description>")
		os.Exit(1)
	}

	taskIDStr := os.Args[2]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Printf("Invalid task ID: %s\n", taskIDStr)
		os.Exit(1)
	}

	newDescription := os.Args[3]

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID %d not found.\n", taskID)
		os.Exit(1)
	}

	err = saveTask(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task with ID %d updated successfully.\n", taskID)
}

func deleteTask() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task-cli delete <task_id>")
		os.Exit(1)
	}

	taskIDStr := os.Args[2]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Printf("Invalid task ID: %s\n", taskIDStr)
		os.Exit(1)
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID %d not found.\n", taskID)
		os.Exit(1)
	}

	err = saveTask(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task with ID %d deleted successfully.\n", taskID)
}

func markTaskStatus(newStatus string) {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: task-cli mark-%s <task_id>\n", newStatus)
		os.Exit(1)
	}

	taskIDStr := os.Args[2]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Printf("Invalid task ID: %s\n", taskIDStr)
		os.Exit(1)
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID %d not found.\n", taskID)
		os.Exit(1)
	}

	err = saveTask(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task with ID %d marked as %s successfully.\n", taskID, newStatus)
}
