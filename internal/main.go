package main

import (
	"log"

	"github.com/xcCodebreaker/Taskmaster/internal/storage"
	"github.com/xcCodebreaker/Taskmaster/internal/task"
	"github.com/xcCodebreaker/Taskmaster/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	taskManager := task.NewTaskManager()

	tasks, err := storage.LoadFromFile("tasks.json")
	if err != nil {
		log.Println("No previous tasks found, starting fresh.")
	} else {
		taskManager.Tasks = tasks
	}

	p := tea.NewProgram(ui.NewModel(taskManager))
	_, err = p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}

	err = storage.SaveToFile(taskManager.Tasks, "tasks.json")
	if err != nil {
		log.Fatalf("Error saving tasks: %v", err)
	}
}
