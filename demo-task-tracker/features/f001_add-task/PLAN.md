# Implementation Plan: Add Task

## Overview
Build the foundation: task model, JSON file storage, and the `add` CLI command. This delivers a working `task-tracker add "description"` command.

## Task Dependencies
Task 1 (model) → Task 2 (storage) → Task 3 (CLI add command)

## Tasks

### Task 1: Task model [PRIORITY: HIGH]
**Corresponds to:** All scenarios (foundation)
**Description:** Define the Task struct with ID, Description, and Done fields.

- [x] Create `internal/task/task.go` with Task struct
- [x] Implement `NewTask(description string) (*Task, error)` with validation
- [x] Write table-driven tests for validation (empty, whitespace, valid)
- [x] Generate ID using `time.UnixNano()` + random suffix (simple, no UUID dep)

**Verification:** `go test ./internal/task/` passes ✓ (4/4 tests)

### Task 2: JSON file storage [PRIORITY: HIGH]
**Corresponds to:** "task should be stored persistently"
**Description:** Implement file-based storage that can load, save, and append tasks.

- [x] Create `internal/store/store.go` with Store struct (wraps file path + tasks slice)
- [x] Implement `Load()` — read JSON file, return existing tasks (empty slice if file doesn't exist)
- [x] Implement `Save()` — write tasks slice to JSON file
- [x] Implement `AddTask(task *Task) error` — append and save
- [ ] Use `io.Reader`/`io.Writer` interfaces for testability (deferred — not needed for current tests)
- [x] Write table-driven tests (empty file, existing tasks, add task)

**Verification:** `go test ./internal/store/` passes ✓ (2/2 tests)

### Task 3: CLI `add` command [PRIORITY: HIGH]
**Corresponds to:** All "add" scenarios
**Description:** Wire up the `add` subcommand that parses args, validates, creates task, saves it.

- [ ] Create `cmd/cli/main.go` with `go run` entry point
- [ ] Parse `os.Args` for `add` subcommand and description argument
- [ ] Handle missing argument with usage message
- [ ] Call `task.NewTask()` → `store.AddTask()` → print confirmation with ID
- [ ] Store config path: `~/.task-tracker/tasks.json` (create dir if needed)
- [ ] Write integration-style test in `cmd/cli/main_test.go`

**Verification:** `go run ./cmd/cli add "Test task"` creates task, prints ID, `tasks.json` exists with correct content

## Technical Decisions
| Decision | Choice | Reason |
|---|---|---|
| ID generation | `time.UnixNano()` + rand suffix | No external deps, sufficient uniqueness for personal use |
| Storage format | JSON array | Human-readable, stdlib encoding/json, simple |
| Config location | `~/.task-tracker/` | Follows XDG-like convention, isolated from other tools |
| CLI parsing | `os.Args` manual parsing | Only 4 commands planned, `flag` subcommands are verbose, no external dep |

## Risks
- None significant — straightforward stdlib-only Go project
