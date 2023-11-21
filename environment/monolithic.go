package environment

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/scheduler/request"
	"example.com/scheduler/task"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
)

type FakeDB struct {
	AdmissionQueue *llq.Queue
	InProcess      map[string]task.TaskV2
}

type MonolithicEnv struct {
	db FakeDB
}

func (e *MonolithicEnv) generateRequest(ch chan<- request.Request, quit <-chan string) {
	for {
		next_arrival := rand.ExpFloat64()
		select {
		case <-time.After(time.Duration(next_arrival * float64(time.Second))):
			ch <- request.GetRequest(5 * rand.Float64())
		case <-quit:
			fmt.Println("End of Request generation.")
			return
		}

	}
}

func (e *MonolithicEnv) AdmissionControl(ch <-chan request.Request, quit <-chan string) {
	for {
		select {
		case r := <-ch:
			fmt.Printf("Admitting new request: %v\n", r)
			e.db.AdmissionQueue.Enqueue(r)
		case <-quit:
			fmt.Println("End of admission control")
			return
		}
	}
}

func (e *MonolithicEnv) WorkerManager(quit <-chan string) {

	// goroutine checking for new incoming requests from admission controller
	ch := make(chan request.Request)
	go func() {
		for {
			// checking for quit signal, but not blocking
			select {
			case <-quit:
				fmt.Println("End of goroutine checking AdmissionQueue")
				return
			default:
			}

			// checking for incoming request
			r, ok := e.db.AdmissionQueue.Dequeue()
			if ok == true {
				ch <- r.(request.Request)
			}
		}
	}()

	man_ch := make(chan string)

	for {
		select {
		// connection to workers
		case m := <-man_ch:
			fmt.Println(m)
			delete(e.db.InProcess, m)
		// connection to goroutine checking for requests
		case r := <-ch:
			t := task.TaskV2{
				Start:       time.Now().Format(time.TimeOnly),
				Computation: time.Duration(r.GetComputation() * float64(time.Second)),
			}
			e.db.InProcess[t.Start] = t

			// running worker
			go func() {
				local_task := t
				fmt.Printf("Starting Task: %v\n", local_task)
				// fail with probability of 0.3
				if rand.Float64() <= 0.3 {
					<-time.After(time.Duration(3 * rand.Float64() * float64(time.Second)))
					man_ch <- fmt.Sprintf("Task %v failed.\n", local_task)
				} else {
					<-time.After(local_task.Computation)
					fmt.Printf("Task %v completed.\n", local_task)
					man_ch <- local_task.Start
				}
				return
			}()
		case <-quit:
			fmt.Println("End of worker manager")
			return
		}

	}

}

func (e *MonolithicEnv) Initialize() {
	e.db.AdmissionQueue = llq.New()
	e.db.InProcess = make(map[string]task.TaskV2)
}

func (e *MonolithicEnv) Run(n int) {
	ch := make(chan request.Request)
	quit := make(chan string)

	go e.generateRequest(ch, quit)
	go e.AdmissionControl(ch, quit)
	go e.WorkerManager(quit)

	timeout := time.After(time.Duration(n * int(time.Second)))

	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Printf("Report ## Queue Length: %v ## Number of Workers: %v\n",
				e.db.AdmissionQueue.Size(), len(e.db.InProcess))
		case <-timeout:
			quit <- "End from Control Plane"
			fmt.Print("End of program")
			<-time.After(500 * time.Millisecond)
			return
		}
	}
}
