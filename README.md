# Simple Task Tracker CLI (https://roadmap.sh/projects/task-tracker)

A lightweight command-line task tracker written in Go. Tasks are persisted locally in a JSON file.

## Project URL

https://github.com/Hla-aung/Simple-Task-Tracker-CLI

## Features

- Add tasks with an optional description
- List all tasks with their status
- Mark tasks as done
- Delete tasks by ID
- Support for multiple JSON task files via a global flag

## Requirements

- Go 1.21+

## Installation

```bash
git clone https://github.com/Hla-aung/Simple-Task-Tracker-CLI.git
cd Simple-Task-Tracker-CLI
go build -o task .
```

## Usage

```
task [command] [flags]
```

### Commands

| Command | Description |
|---------|-------------|
| `add`   | Add a new task |
| `list`  | List all tasks |
| `done`  | Mark a task as done |
| `delete`| Delete a task |

### Global Flags

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--file` | `-f` | `tasks.json` | Path to the JSON file used to store tasks |

### Add a task

```bash
task add "Buy groceries"
task add "Buy groceries" -d "Milk, eggs, bread"
```

### List tasks

```bash
task list
```

Output example:
```
1. Buy groceries [ ]
   Milk, eggs, bread
2. Write report [x]
```

### Mark a task as done

```bash
task done 1
```

### Delete a task

```bash
task delete 1
```

### Use a custom task file

```bash
task -f work.json add "Fix bug #42"
task -f work.json list
```

## Data Storage

Tasks are stored as a JSON array in `tasks.json` (or the file specified with `-f`). The file is created automatically on the first `add` command.

## License

MIT
