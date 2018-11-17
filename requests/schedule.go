package requests

import (
	"time"
)

type Schedule interface {
	Run()
}

type RequestSchedule struct {
	interval      time.Duration
	count         int
	request       Request
	maxConcurrent int
}

func NewSchedule(interval time.Duration, count int, request Request, maxConcurrent int) (*RequestSchedule) {
	return &RequestSchedule{
		interval:      interval,
		count:         count,
		request:       request,
		maxConcurrent: maxConcurrent,
	}
}

// We are going to run our request `count` amount of times on the
// `interval` using `maxConcurrent` amount of go routines
func (r RequestSchedule) Run() {
	go func() {
		limitChannel := make(chan bool, r.maxConcurrent)
		for i := 0; i < r.count; i++ {

			// this is our limiter
			// this channel is buffered so it will not block until
			// you have placed maxConcurrent items without any being removed
			// an item will be removed from the channel when the request is complete
			limitChannel <- true

			go func() {
				r.request.Run()
				<-limitChannel
			}()
		}

	}()
}
