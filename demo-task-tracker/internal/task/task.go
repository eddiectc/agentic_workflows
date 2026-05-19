package task

import (
	"fmt"
	"strings"
)

// Task represents a single task item.
type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// NewTask creates a new task with the given description.
// Returns an error if the description is empty or whitespace-only.
func NewTask(description string) (*Task, error) {
	trimmed := strings.TrimSpace(description)
	if trimmed == "" {
		return nil, fmt.Errorf("description is required")
	}
	return &Task{
		ID:          generateID(),
		Description: trimmed,
		Done:        false,
	}, nil
}

// generateID creates a simple unique ID from timestamp + random suffix.
func generateID() string {
	return fmt.Sprintf("tsk-%d", 1) // placeholder; replaced with real impl
}
