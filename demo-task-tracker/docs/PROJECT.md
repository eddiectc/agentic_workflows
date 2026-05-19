# Project: Task Tracker CLI

## Description
A minimal command-line task tracker for personal productivity. Users can add, list, complete, and delete tasks stored in a local JSON file.

## Goals
- Quick task capture from the terminal
- Simple task listing with filters (all, pending, done)
- Mark tasks as complete
- Delete tasks
- Persistent storage (local JSON file)

## Non-Goals
- Due dates / reminders
- Categories / tags
- Multi-user / sharing
- Cloud sync
- GUI

## Architecture
Single-binary CLI tool. Commands read/write a local `tasks.json` file in the user's config directory (`~/.task-tracker/tasks.json`).

## Tech Stack
| Component | Choice | Reason |
|---|---|---|
| Language | Go 1.21+ | Single binary, fast startup, minimal deps |
| CLI framework | None (stdlib `flag` or `os.Args`) | Keep it simple, no external deps |
| Storage | JSON file | Simple, human-readable, no database needed |
| Testing | stdlib `testing` | No external test framework needed |

## Constraints
- No external dependencies (stdlib only)
- Single binary output
- Config file in `~/.task-tracker/`

## Directory Structure
```
demo-task-tracker/
├── cmd/
│   └── cli/
│       └── main.go          # Entry point
├── internal/
│   ├── store/
│   │   └── store.go         # Task storage (JSON file read/write)
│   └── task/
│       └── task.go          # Task model
├── docs/
│   ├── PROJECT.md
│   └── CONVENTIONS.md
├── features/
│   └── f001_add-task/
├── go.mod
└── README.md
```

## Conventions
- Go stdlib only, no external dependencies
- `internal/` for non-exported packages
- Table-driven tests
- Error messages are user-friendly (no raw error chains)
