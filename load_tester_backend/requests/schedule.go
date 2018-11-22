package requests

import (
	"time"
)

type Schedule struct {
	Id              string
	Name            string
	StartTime       time.Time
	EndTime         time.Time
	RequestCount    int
	AverageDuration int
	ErrorCount      int
}

type RequestSchedule struct {
	interval        time.Duration
	count           int
	request         *Request
	maxConcurrent   int
	errorChannel    chan error
	responseChannel chan RequestResponse
}

func NewSchedule(
	interval time.Duration,
	count int,
	request *Request,
	maxConcurrent int) (*RequestSchedule) {
	return &RequestSchedule{
		interval:      interval,
		count:         count,
		request:       request,
		maxConcurrent: maxConcurrent,
	}
}

// We are going to run our request `count` amount of times on the
// `interval` using `maxConcurrent` amount of go routines
// this method returns a channel where it will write 1 and only 1 message slice,
// 1 and only one duration slice, one and only one error slice (hopefully nil)
func (r RequestSchedule) Run() (chan RequestResponse, chan bool) {
	limitChannel := make(chan bool, r.maxConcurrent)
	responseChannel := make(chan RequestResponse)
	doneChannel := make(chan bool)
	responseIndex := 0
	// the `requests` message, duration, and error channel will always be overridden
	go func() {
		for i := 0; i < r.count; i++ {
			limitChannel <- true
			go func() {
				msg, dur, err := r.request.Run()
				errMsg := ""
				if err != nil {
					errMsg = err.Error()
				}
				resp := RequestResponse{
					Message:     msg,
					ErrString:   errMsg,
					Duration:    dur,
					RequestUrl:  r.request.RequestUrl,
					RequestType: r.request.RequestType,
				}

				responseChannel <- resp
				// if we've written the same amount of responses to our channel as we wanted to
				// or `r.count` we know we are done so write one result to done channel
				if responseIndex == (r.count - 1) {
					doneChannel <- true
				} else {
					responseIndex++
				}
				<-limitChannel
			}()
			// this is our limiter
			// this channel is buffered so it will not block until
			// you have placed maxConcurrent items without any being removed
			// an item will be removed from the channel when the request is complete
			requestDuration := r.interval
			select {
			case <-time.After(requestDuration):
				continue
			}
		}
	}()
	return responseChannel, doneChannel
}
