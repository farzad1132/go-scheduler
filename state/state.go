package state

import (
	"example.com/scheduler/list"
	"example.com/scheduler/request"
	"example.com/scheduler/task"
)

type State struct {
	status          bool
	number_in_queue int
	system_queue    list.PriorityQueue
	server          list.PriorityQueue
}

func (s *State) Initialize() {
	s.status = false
	s.number_in_queue = 0
	//s.system_queue = list.PriorityQueue{}
	s.system_queue.Initialize()

	s.server.Initialize()
}

func (s *State) GetStatus() bool {
	return s.status
}

func (s *State) GetNumberInQueue() int {
	return s.number_in_queue
}

func (s *State) AddRequest(r request.Request) {
	s.system_queue.Add(r)
	s.number_in_queue = s.system_queue.Len()
}

func (s *State) RemoveRequest() request.Request {
	r, _ := s.system_queue.Delete()
	s.number_in_queue = s.system_queue.Len()
	return r.(request.Request)
}

func (s *State) AddToServer(t task.Task) {
	s.server.Add(t)
	s.status = true
}

func (s *State) RemoveFromServer() task.Task {
	t, _ := s.server.Delete()
	if s.server.Len() == 0 {
		s.status = false
	}
	return t.(task.Task)
}

func (s *State) TimeToFinish(t int) float64 {
	item, _ := s.server.Peek()
	return item.(list.HasPriority).GetPriority() - float64(t)
}
