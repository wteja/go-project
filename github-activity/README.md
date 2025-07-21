## GitHub Activity CLI

This is a simple command-line application written in Go that fetches and summarizes recent GitHub activity for a given user. It groups activities such as pushes, pull requests, issues, releases, creations, comments, and watches by repository.

### Features
- Fetches recent public events for a GitHub user
- Groups activities by repository and event type
- Displays a summary of actions (pushes, pull requests, issues, releases, etc.)

### Prerequisites
- Go installed (for building from source)
- Internet connection

### Usage

#### Download/Build
If you have the binary (`github-activity.exe`), you can run it directly. To build from source:

```sh
go build -o github-activity.exe main.go github.go
```

#### Run

```sh
./github-activity.exe <github-username>
```

Replace `<github-username>` with the GitHub username you want to check.

#### Example

```sh
./github-activity.exe octocat
```

### Output
The app will print a summary of recent activities grouped by repository, for example:

```
Pushed 3 commits to octocat/Hello-World
Opened 1 pull requests in octocat/Hello-World
Opened 2 issues in octocat/Hello-World
Released 1 versions in octocat/Hello-World
Created 1 items in octocat/Hello-World
Commented on 2 issues in octocat/Hello-World
Watched 1 times in octocat/Hello-World
```

### Error Handling
If the username is invalid or there is a network/API error, an error message will be displayed.

### License
See [LICENSE](../LICENSE) for details.
