package list

import (
	"slices"

	"example.com/scheduler/task"
)

type SimpleList struct {
	container []task.Task
}

func (l *SimpleList) Initialize() {
	l.container = make([]task.Task, 0)
}

func (l *SimpleList) AddTask(task task.Task) {
	l.container = append(l.container, task)
}

func (l *SimpleList) RemoveTask() task.Task {
	task := l.container[0]
	l.container = slices.Delete(l.container, 0, 1)
	return task
}
