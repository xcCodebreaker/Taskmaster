package ui

import (
	"fmt"
	"strings"

	"github.com/xcCodebreaker/Taskmaster/internal/task"

	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	normalMode mode = iota
	addTaskMode
)

type model struct {
	tasks    *task.TaskManager
	selected int
	mode     mode
	input    string
}

func NewModel(tasks *task.TaskManager) model {
	return model{tasks: tasks, selected: 0, mode: normalMode, input: ""}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.mode {
		case normalMode:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "a":
				m.mode = addTaskMode
				m.input = ""
			case "d":
				if len(m.tasks.ListTasks()) > m.selected {
					m.tasks.DeleteTask(m.tasks.ListTasks()[m.selected].ID)
				}
			case "m":
				if len(m.tasks.ListTasks()) > m.selected {
					m.tasks.ToggleTaskDone(m.tasks.ListTasks()[m.selected].ID)
				}
			case "u":
				m.tasks.Undo()
			case "r":
				m.tasks.Redo()
			case "down", "j":
				if m.selected < len(m.tasks.ListTasks())-1 {
					m.selected++
				}
			case "up", "k":
				if m.selected > 0 {
					m.selected--
				}
			}
		case addTaskMode:
			switch msg.String() {
			case "enter":
				if strings.TrimSpace(m.input) != "" {
					m.tasks.AddTask(m.input)
				}
				m.mode = normalMode
			case "esc":
				m.mode = normalMode
			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			default:
				m.input += msg.String()
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	if m.mode == addTaskMode {
		return fmt.Sprintf("Enter task name: %s\n[Enter] to confirm • [Esc] to cancel", m.input)
	}

	s := "Tasks:\n"
	for i, t := range m.tasks.ListTasks() {
		cursor := " "
		if i == m.selected {
			cursor = ">"
		}
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		s += fmt.Sprintf("%s %s %s\n", cursor, status, t.Description)
	}
	s += "\n[↑/↓] Navigate • [a] Add Task • [d] Delete Task • [m] Toggle Done • [u] Undo • [r] Redo • [q] Quit\n"
	return s
}
