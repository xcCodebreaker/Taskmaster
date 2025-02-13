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
