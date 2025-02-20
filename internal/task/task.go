package task

type Task struct {
	ID          int
	Description string
	Done        bool
}

type TaskManager struct {
	tasks       []Task
	history     [][]Task
	redoHistory [][]Task
}

// Find out about constructor functions
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: []Task{},
	}
}

// The method name speaks for itself about what it does...
func (tm *TaskManager) getNextID() int {
	if len(tm.tasks) == 0 {
		return 1
	}

	return tm.tasks[len(tm.tasks)-1].ID + 1
}

// Necessary for undo/redo functions. Any changes made in the program, this func will be run everytime.
func (tm *TaskManager) saveState() {
	state := make([]Task, len(tm.tasks))
	copy(state, tm.tasks)
	tm.history = append(tm.history, state)
	tm.redoHistory = nil
}

func (tm *TaskManager) AddTask(description string) {
	tm.saveState()
	tm.tasks = append(tm.tasks, Task{
		ID:          tm.getNextID(),
		Description: description,
		Done:        false,
	})
}

func (tm *TaskManager) DeleteTask(id int) {
	tm.saveState()
	for i, t := range tm.tasks {
		if t.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
		}
	}
}

func (tm *TaskManager) ToggleTaskDone(id int) {
	for i, t := range tm.tasks {
		if t.ID == id {
			tm.saveState()
			tm.tasks[i].Done = !t.Done
			return
		}
	}
}
