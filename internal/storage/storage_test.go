package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/xcCodebreaker/Taskmaster/internal/task"
)

func sampleTasks() []task.Task {
	return []task.Task{
		{ID: 1, Description: "Task 1", Done: false},
		{ID: 2, Description: "Task 2", Done: true},
	}
}

func TestSaveToFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "tasks.json")

	tasks := sampleTasks()

	if err := SaveToFile(tasks, filePath); err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	var loadedTasks []task.Task
	if err = json.Unmarshal(data, &loadedTasks); err != nil {
		t.Fatalf("failed to unmarshal JSON data: %v", err)
	}

	if len(loadedTasks) != len(tasks) {
		t.Fatalf("expected %d tasks, got %d", len(tasks), len(loadedTasks))
	}
	for i, tsk := range tasks {
		if loadedTasks[i] != tsk {
			t.Errorf("expected task %v, got %v", tsk, loadedTasks[i])
		}
	}
}
