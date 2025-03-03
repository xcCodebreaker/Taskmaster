package task

type Task struct {
	ID          int
	Description string
	Done        bool
}

type TaskManager struct {
	Tasks       []Task
	History     [][]Task
	RedoHistory [][]Task
}

// Find out about constructor functions
func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: []Task{},
	}
}

// The method name speaks for itself about what it does...
func (tm *TaskManager) getNextID() int {
	if len(tm.Tasks) == 0 {
		return 1
	}

	return tm.Tasks[len(tm.Tasks)-1].ID + 1
}

// Necessary for undo/redo functions. Any changes made in the program, this func will be run everytime.
func (tm *TaskManager) saveState() {
	state := make([]Task, len(tm.Tasks))
	copy(state, tm.Tasks)
	tm.History = append(tm.History, state)
	tm.RedoHistory = nil
}

func (tm *TaskManager) AddTask(description string) {
	tm.saveState()
	tm.Tasks = append(tm.Tasks, Task{
		ID:          tm.getNextID(),
		Description: description,
		Done:        false,
	})
}

func (tm *TaskManager) DeleteTask(id int) {
	tm.saveState()
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
		}
	}
}

func (tm *TaskManager) ToggleTaskDone(id int) {
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.saveState()
			tm.Tasks[i].Done = !t.Done
			return
		}
	}
}

func (tm *TaskManager) Undo() error {
	if len(tm.History) == 0 {
		return nil
	}
	tm.RedoHistory = append(tm.RedoHistory, tm.Tasks)
	tm.Tasks = tm.History[len(tm.History)-1]
	tm.History = tm.History[:len(tm.History)-1]
	return nil
}

func (tm *TaskManager) Redo() error {
	if len(tm.RedoHistory) == 0 {
		return nil
	}
	tm.History = append(tm.History, tm.Tasks)
	tm.Tasks = tm.RedoHistory[len(tm.RedoHistory)-1]
	tm.RedoHistory = tm.RedoHistory[:len(tm.RedoHistory)-1]
	return nil
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}
