# CLI To-Do List

A simple, yet powerful command-line to-do list application built in Go. Manage tasks with priorities, notes, and persistent storage—all from your terminal.

## Features

1. **Add Tasks**: Create tasks with customizable priorities (high, medium, low) and optional notes for extra context.
2. **Priority Sorting**: List tasks sorted by priority (high > medium > low) with the -sort flag.
3. **Task Management**: Edit titles, toggle completion status, and delete tasks by index.
4. **Notes**: Attach short descriptions or reminders to tasks.
5. **Persistence**: Tasks are saved to todos.json and loaded automatically on startup.

## Installation

1. **Prerequisites**: Ensure you have `Go` installed (version 1.16+ recommended).
2. **Clone the Repository**

```bash
git clone https://github.com/Prateek2593/todo-cli.git
cd todo-cli
```

3. **Build the Application**

```bash
go build -o todo-cli
```

4. **Run it**

```bash
./todo-cli
```

## Usage

Run the app with various flags to manage your to-do list:

### Add a task

```bash
./todo-cli -add "Finish report" -p high -n "Due by EOD"
```

- `-add` - Add a new task.
- `-p` - Priority (`high`, `medium`, `low`; default to `medium`).
- `-n` - Optional notes for the task.

### List tasks

```bash
./todo-cli -list
```

- Add `-sort` to sort tasks by priority.

```bash
./todo-cli -list -sort
```

### Edit a task

```bash
./todo-cli -edit "0:Write code"
```

- Format - `index:new_title`

### Toggle completion

```bash
./todo-cli -toggle 0
```

- Marks task at index `0` as completed.

### Delete a task

```bash
./todo-cli -del 0
```

- Removes task at index `0`.

## Project Structure

- `main.go`: Entry point for the application.
- `todo.go`: Contains the `Todo` struct and methods for managing tasks.
- `command.go`: Defines command-line flags and their handling.
- `storage.go`: Handles saving and loading tasks from `todos.json`.

## Dependencies

- github.com/aquasecurity/table: For rendering the task table.
- golang.org/x/text: For title-casing priorities.

Install them with:

```go
go get github.com/aquasecurity/table
go get golang.org/x/text/cases
go get golang.org/x/text/language
```

## What I Learned

- **Go Fundamentals**: Gained hands-on experience with structs, slices, and the standard library (e.g., flag, json).
- **CLI Development**: Designed a user-friendly interface with flexible command-line options.
- **Debugging**: Resolved issues like sorting discrepancies and data persistence through iterative testing.
- **Code Organization**: Applied modular design principles to keep the codebase clean and maintainable.
