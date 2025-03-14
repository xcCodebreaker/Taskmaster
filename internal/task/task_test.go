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

func TestDeleteTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task 1")
	tm.AddTask("Task 2")

	tm.DeleteTask(1)
	if len(tm.Tasks) != 1 {
		t.Errorf("expected 1 task after deletion, got %d", len(tm.Tasks))
	}
	if tm.Tasks[0].ID != 2 {
		t.Errorf("expected remaining task ID to be 2, got %d", tm.Tasks[0].ID)
	}
}

func TestToggleTaskDone(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Toggle Task")

	if tm.Tasks[0].Done {
		t.Error("expected task to be not done initially")
	}

	tm.ToggleTaskDone(1)
	if !tm.Tasks[0].Done {
		t.Error("expected task to be done after toggling")
	}

	tm.ToggleTaskDone(1)
	if tm.Tasks[0].Done {
		t.Error("expected task to be not done after toggling again")
	}

}
