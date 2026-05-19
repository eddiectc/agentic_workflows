# Feature: Add Task

## Description
Users can add a new task with a description. The task is stored persistently and assigned a unique ID.

## User Stories
- As a user, I want to add a task with a description, so that I can remember to do it later
- As a user, I want confirmation that my task was added, so that I know it was saved

## Scenarios

### Scenario: Add a task with a simple description
**Given** I have no existing tasks
**When** I run `task-tracker add "Buy groceries"`
**Then** I should see confirmation that the task was added
**And** The task should be assigned a unique ID
**And** The task should be stored persistently

### Scenario: Add multiple tasks
**Given** I have one existing task
**When** I run `task-tracker add "Call dentist"`
**Then** A new task should be created with a different ID
**And** The existing task should remain unchanged

### Scenario: Add a task with an empty description
**Given** Any state
**When** I run `task-tracker add ""`
**Then** I should see an error message saying description is required
**And** No task should be created

### Scenario: Add a task with whitespace-only description
**Given** Any state
**When** I run `task-tracker add "   "`
**Then** I should see an error message saying description is required
**And** No task should be created

## Edge Cases
- Very long descriptions (should be accepted, no arbitrary limit)
- Descriptions with special characters (quotes, newlines, unicode)
- Running add command with no arguments at all

## Constraints
- Description must be non-empty and non-whitespace
- Each task gets a unique, auto-assigned ID
- Tasks are stored locally (no network)

## Non-Goals
- Due dates
- Priority levels
- Categories / tags
- Task editing (separate feature)

## Dependencies
- None (this is the first feature)
