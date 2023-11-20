package environment

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/scheduler/request"
	"example.com/scheduler/state"
	"example.com/scheduler/task"
)

type SimpleEnv struct {
	state state.State
}

func (e *SimpleEnv) Initialize() {
	e.state.Initialize()
}

func (e *SimpleEnv) generateRequest() (request.Request, bool) {
	if rand.Float64() <= 0.4 {
		return request.GetRequest(5 * rand.Float64()), true
	} else {
		return request.Request{}, false
	}
}

func (e *SimpleEnv) AddTaskToServer(t int) {
	if e.state.GetNumberInQueue() > 0 {
		r := e.state.RemoveRequest()
		job := task.Task{
			Name:     fmt.Sprintf("Task at %v", t),
			Priority: float64(t) + r.GetComputation(),
		}
		e.state.AddToServer(job)
		fmt.Printf("Staring the process of task: %v\n", job)
	}
}

func (e *SimpleEnv) Run(n int) {
	for t := 0; t < n; t++ {
		fmt.Printf("System at time: %v, Number in queue: %d, Server Status: %v\n",
			t, e.state.GetNumberInQueue(), e.state.GetStatus())

		r, ok := e.generateRequest()
		if ok == true {
			fmt.Printf("Admitting new Request %v to the system.\n", r)
			e.state.AddRequest(r)
		}

		if e.state.GetStatus() == false {
			e.AddTaskToServer(t)
		} else {
			if e.state.TimeToFinish(t) <= 0 {
				job := e.state.RemoveFromServer()
				fmt.Printf("Task %v completed.\n", job)
				e.AddTaskToServer(t)
			}
		}

		time.Sleep(time.Second)

	}
}
