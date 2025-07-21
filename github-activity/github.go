package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubEventError struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
	Status           string `json:"status"`
}

type GitHubEvent struct {
	ID   string                `json:"id"`
	Type string                `json:"type"`
	Repo GitHubEventRepository `json:"repo"`
}

type GitHubEventRepository struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getGitHubActivity(username string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp GitHubEventError
		_ = json.Unmarshal(bytes, &errResp)
		return nil, fmt.Errorf("GitHub API error: %s (status: %s)", errResp.Message, errResp.Status)
	}

	var activities []string
	var events []GitHubEvent
	err = json.Unmarshal(bytes, &events)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	pushEvents := make([]GitHubEvent, 0)
	pullRequestEvents := make([]GitHubEvent, 0)
	issueEvents := make([]GitHubEvent, 0)
	releaseEvents := make([]GitHubEvent, 0)
	createEvents := make([]GitHubEvent, 0)
	issueCommentEvents := make([]GitHubEvent, 0)
	watchEvents := make([]GitHubEvent, 0)

	groupPushByRepo := make(map[string]int)
	groupPullRequestByRepo := make(map[string]int)
	groupIssueByRepo := make(map[string]int)
	groupReleaseByRepo := make(map[string]int)
	groupCreateByRepo := make(map[string]int)
	groupIssueCommentByRepo := make(map[string]int)
	groupWatchByRepo := make(map[string]int)

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			pushEvents = append(pushEvents, event)
		case "PullRequestEvent":
			pullRequestEvents = append(pullRequestEvents, event)
		case "IssuesEvent":
			issueEvents = append(issueEvents, event)
		case "ReleaseEvent":
			releaseEvents = append(releaseEvents, event)
		case "CreateEvent":
			createEvents = append(createEvents, event)
		case "IssueCommentEvent":
			issueCommentEvents = append(issueCommentEvents, event)
		case "WatchEvent":
			watchEvents = append(watchEvents, event)
		}
	}

	// Grouping events by repository
	for _, event := range pushEvents {
		groupPushByRepo[event.Repo.Name]++
	}
	for repo, count := range groupPushByRepo {
		activities = append(activities, fmt.Sprintf("Pushed %d commits to %s", count, repo))
	}

	// Grouping pull request events by repository
	for _, event := range pullRequestEvents {
		groupPullRequestByRepo[event.Repo.Name]++
	}
	for repo, count := range groupPullRequestByRepo {
		activities = append(activities, fmt.Sprintf("Opened %d pull requests in %s", count, repo))
	}

	// Grouping issue events by repository
	for _, event := range issueEvents {
		groupIssueByRepo[event.Repo.Name]++
	}
	for repo, count := range groupIssueByRepo {
		activities = append(activities, fmt.Sprintf("Opened %d issues in %s", count, repo))
	}

	// Grouping release events by repository
	for _, event := range releaseEvents {
		groupReleaseByRepo[event.Repo.Name]++
	}
	for repo, count := range groupReleaseByRepo {
		activities = append(activities, fmt.Sprintf("Released %d versions in %s", count, repo))
	}

	// Grouping create events by repository
	for _, event := range createEvents {
		groupCreateByRepo[event.Repo.Name]++
	}
	for repo, count := range groupCreateByRepo {
		activities = append(activities, fmt.Sprintf("Created %d items in %s", count, repo))
	}

	// Grouping issue comment events by repository
	for _, event := range issueCommentEvents {
		groupIssueCommentByRepo[event.Repo.Name]++
	}
	for repo, count := range groupIssueCommentByRepo {
		activities = append(activities, fmt.Sprintf("Commented on %d issues in %s", count, repo))
	}

	// Grouping watch events by repository
	for _, event := range watchEvents {
		groupWatchByRepo[event.Repo.Name]++
	}
	for repo, count := range groupWatchByRepo {
		activities = append(activities, fmt.Sprintf("Watched %d times in %s", count, repo))
	}

	return activities, nil
}
