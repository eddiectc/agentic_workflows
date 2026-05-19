package store

import (
	"os"
	"path/filepath"
	"testing"

	"task-tracker/internal/task"
)

func TestStore_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	store := New(path)

	// Load from non-existent file should return empty
	if err := store.Load(); err != nil {
		t.Fatal(err)
	}
	if got := len(store.Tasks()); got != 0 {
		t.Errorf("expected 0 tasks, got %d", got)
	}

	// Add a task
	tk, err := task.NewTask("Test task")
	if err != nil {
		t.Fatal(err)
	}
	if err := store.AddTask(tk); err != nil {
		t.Fatal(err)
	}

	// Reload and verify
	store2 := New(path)
	if err := store2.Load(); err != nil {
		t.Fatal(err)
	}
	tasks := store2.Tasks()
	if got := len(tasks); got != 1 {
		t.Fatalf("expected 1 task, got %d", got)
	}
	if tasks[0].Description != "Test task" {
		t.Errorf("expected 'Test task', got %q", tasks[0].Description)
	}
}

func TestStore_ExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")

	// Pre-populate file
	existing := `[{"id":"1","description":"Existing","done":false}]`
	if err := os.WriteFile(path, []byte(existing), 0o644); err != nil {
		t.Fatal(err)
	}

	store := New(path)
	if err := store.Load(); err != nil {
		t.Fatal(err)
	}
	if got := len(store.Tasks()); got != 1 {
		t.Fatalf("expected 1 task, got %d", got)
	}
}
