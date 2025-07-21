package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := os.Args[1]
	activities, err := getGitHubActivity(username)
	if err != nil {
		fmt.Printf("Error fetching GitHub activity: %v\n", err)
		return
	}

	for _, activity := range activities {
		fmt.Println(activity)
	}
}
