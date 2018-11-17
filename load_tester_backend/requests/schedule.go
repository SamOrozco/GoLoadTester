package requests

import (
	"time"
)

type Schedule interface {
	Run()
}

type RequestSchedule struct {
	interval       time.Duration
	count          int
	request        *Request
	maxConcurrent  int
	messageResult  chan []string
	durationResult chan []int
	errorResult    chan []error
}

func NewSchedule(
	interval time.Duration,
	count int,
	request *Request,
	maxConcurrent int) (*RequestSchedule) {
	return &RequestSchedule{
		interval:       interval,
		count:          count,
		request:        request,
		maxConcurrent:  maxConcurrent,
		messageResult:  make(chan []string),
		durationResult: make(chan []int),
		errorResult:    make(chan []error),
	}
}

// We are going to run our request `count` amount of times on the
// `interval` using `maxConcurrent` amount of go routines
// this method returns a channel where it will write 1 and only 1 message slice,
// 1 and only one duration slice, one and only one error slice (hopefully nil)
func (r RequestSchedule) Run() (*ScheduleResponse) {

	limitChannel := make(chan bool, r.maxConcurrent)
	msgChannel := make(chan string)
	messages := make([]string, 0)
	durChannel := make(chan int)
	durations := make([]int, 0)
	errors := make([]error, 0)
	errorChannel := make(chan error)

	// the `requests` message, duration, and error channel will always be overridden
	go func() {
		for i := 0; i < r.count; i++ {
			limitChannel <- true
			go func() {
				msg, dur, err := r.request.Run()
				<-limitChannel
				msgChannel <- msg
				durChannel <- dur
				errorChannel <- err
			}()
			// this is our limiter
			// this channel is buffered so it will not block until
			// you have placed maxConcurrent items without any being removed
			// an item will be removed from the channel when the request is complete
			<-time.After(r.interval)
		}
	}()

	go func() {
		for i := *new(int); i < r.count; i++ {
			messages = append(messages, <-msgChannel)
			durations = append(durations, <-durChannel)
			errors = append(errors, <-errorChannel)
		}
		r.messageResult <- messages
		r.durationResult <- durations
		r.errorResult <- errors
	}()

	respone := &ScheduleResponse{
		MessageResponse:  r.messageResult,
		DurationResponse: r.durationResult,
		ErrorResponse:    r.errorResult,
	}

	return respone
}
