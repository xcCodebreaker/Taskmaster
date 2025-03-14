package task

import (
	"testing"
)

func TestNewTaskManager(t *testing.T) {
	tm := NewTaskManager()
	if tm == nil {
		t.Fatal("NewTaskManager returned nil")
	}
	if len(tm.Tasks) != 0 {
		t.Errorf("expected empty task list, got %d tasks", len(tm.Tasks))
	}
}

func TestAddTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Test Task")
	if len(tm.Tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tm.Tasks))
	}
	task := tm.Tasks[0]
	if task.ID != 1 {
		t.Errorf("expected task ID 1, got %d", task.ID)
	}
	if task.Description != "Test Task" {
		t.Errorf("expected description 'Test Task', got '%s'", task.Description)
	}
	if task.Done {
		t.Error("newly added task should not be marked as done")
	}
}
