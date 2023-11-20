package task

import "fmt"

type Task struct {
	Name     string
	Priority float64
}

func (t Task) GetPriority() float64 {
	return t.Priority
}

func (t Task) String() string {
	return fmt.Sprintf("[N: %v, P: %v]", t.Name, t.Priority)
}
