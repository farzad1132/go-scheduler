package request

import "fmt"

type Request struct {
	computation float64
	priority    float64
}

func (r Request) GetPriority() float64 {
	return r.priority
}

func (r Request) String() string {
	return fmt.Sprintf("[C: %v, P: %v]", r.computation, r.priority)
}

func GetRequest(c float64) Request {
	return Request{
		computation: c,
	}
}

func (r *Request) GetComputation() float64 {
	return r.computation
}
