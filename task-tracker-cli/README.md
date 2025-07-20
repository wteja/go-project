# Task Tracker CLI

A simple command-line tool for managing tasks, written in Go.
Project URL: https://github.com/wteja/go-project/tree/main/task-tracker-cli

## Features
- Add, update, delete tasks
- Mark tasks as todo, in-progress, or done
- List all tasks
- Persistent storage in `tasks.json`

## Usage

```sh
task-tracker-cli <command> [arguments]
```

### Available Commands

- `add <task_description>`: Add a new task with the given description
- `list`: List all tasks
- `update <task_id> <new_description>`: Update the description of a task
- `delete <task_id>`: Delete a task by its ID
- `mark-todo <task_id>`: Mark a task as "todo"
- `mark-in-progress <task_id>`: Mark a task as "in-progress"
- `mark-done <task_id>`: Mark a task as "done"

## Examples

Add a new task:
```sh
task-tracker-cli add "Write documentation"
```

List all tasks:
```sh
task-tracker-cli list
```

Update a task:
```sh
task-tracker-cli update 1 "Write detailed documentation"
```

Delete a task:
```sh
task-tracker-cli delete 1
```

Mark a task as in progress:
```sh
task-tracker-cli mark-in-progress 1
```

Mark a task as done:
```sh
task-tracker-cli mark-done 1
```

Mark a task as done (revert):
```sh
task-tracker-cli mark-todo 1
```

## Data Storage
Tasks are stored in a local `tasks.json` file in the project directory.

## Build

To build the CLI tool:
```sh
go build -o task-tracker-cli
```

## License
See [LICENSE](../LICENSE).
