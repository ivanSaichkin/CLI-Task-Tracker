# CLI-Task-Tracker

A simple and efficient command-line task manager built with Go.

## 🚀 Features

- ✅ Add new tasks
- 📋 List all tasks
- ✔️ Mark tasks as completed
- 🗑️ Remove tasks
- 🔍 Filter tasks by status (pending/completed)
- 💾 JSON file storage

## 📦 Installation

```bash
go build -o TaskTracker cmd/main.go
```

## 🛠️ Usage

```bash
./TaskTracker list | add <description> | done <id> | get <id> | remove <id> | update <id> <new description> | completed | pending
```

## 🏗️ Project Architecture

```
CLI-Task-Tracker/
├── cmd/
│   └── main.go         # Entry point
├── internal/
│   ├── CLI/
│   │   └── CLI.go          # Command line handling
│   ├── Storage/
│   │   ├── Storage.go      # Storage interface
│   └── Task/
│       └── Task.go         # Task model
├── go.mod
├── go.sum
└── tasks.json              # Data file (auto-created)
```
