package storage

import (
	"encoding/json"
	"os"

	"github.com/xcCodebreaker/Taskmaster/internal/task"
)

func SaveToFile(tasks []task.Task, filename string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadFromFile(filename string) ([]task.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tasks []task.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
