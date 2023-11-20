package task

import (
	"fmt"
	"time"
)

type TaskV2 struct {
	Start       string
	Computation time.Duration
}

func (t TaskV2) String() string {
	return fmt.Sprintf("[S: %v, C: %v]", t.Start, t.Computation)
}
