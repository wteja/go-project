package main

import (
	"fmt"
	"os"
)

const (
	tasksFile        = "tasks.json"
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		fmt.Println("Available commands: add, list, update, delete, mark-todo, mark-in-progress, mark-donw, list")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask()
	case "update":
		updateTask()
	case "delete":
		deleteTask()
	case "mark-todo":
		markTaskStatus(StatusTodo)
	case "mark-in-progress":
		markTaskStatus(StatusInProgress)
	case "mark-done":
		markTaskStatus(StatusDone)
	case "list":
		listTasks()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
