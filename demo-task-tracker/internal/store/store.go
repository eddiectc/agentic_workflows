package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"task-tracker/internal/task"
)

// Store manages persistent task storage in a JSON file.
type Store struct {
	path   string
	tasks  []*task.Task
}

// New creates a new Store for the given file path.
func New(path string) *Store {
	return &Store{path: path}
}

// Load reads tasks from the JSON file.
// Returns empty slice if the file doesn't exist.
func (s *Store) Load() error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			s.tasks = []*task.Task{}
			return nil
		}
		return fmt.Errorf("read tasks file: %w", err)
	}

	if len(data) == 0 {
		s.tasks = []*task.Task{}
		return nil
	}

	var tasks []*task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("parse tasks file: %w", err)
	}
	s.tasks = tasks
	return nil
}

// Save writes all tasks to the JSON file.
// Creates parent directories if they don't exist.
func (s *Store) Save() error {
	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create store directory: %w", err)
	}

	data, err := json.MarshalIndent(s.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}

	if err := os.WriteFile(s.path, data, 0o644); err != nil {
		return fmt.Errorf("write tasks file: %w", err)
	}
	return nil
}

// AddTask appends a task and persists to file.
func (s *Store) AddTask(t *task.Task) error {
	s.tasks = append(s.tasks, t)
	return s.Save()
}

// Tasks returns all loaded tasks.
func (s *Store) Tasks() []*task.Task {
	return s.tasks
}
