package list

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type PriorityQueue struct {
	queue pq.Queue
}

type HasPriority interface {
	GetPriority() float64
}

func (q *PriorityQueue) Initialize() {

	q.queue = *pq.NewWith(func(a, b interface{}) int {
		prioA := a.(HasPriority).GetPriority()
		prioB := b.(HasPriority).GetPriority()
		return int(utils.Float64Comparator(prioA, prioB))
	})
}

func (q *PriorityQueue) Add(task interface{}) {
	q.queue.Enqueue(task)
}

func (q *PriorityQueue) Delete() (interface{}, bool) {
	return q.queue.Dequeue()
}

func (q *PriorityQueue) Len() int {
	return q.queue.Size()
}

func (q PriorityQueue) String() string {
	return q.queue.String()
}

func (q *PriorityQueue) Peek() (interface{}, bool) {
	return q.queue.Peek()
}
