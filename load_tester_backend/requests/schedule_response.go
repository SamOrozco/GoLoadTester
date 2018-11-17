package requests

type ScheduleResponse struct {
	MessageResponse  chan []string
	DurationResponse chan []int
	ErrorResponse    chan []error
}

func (s ScheduleResponse) getMessages() ([]string) {
	return <-s.MessageResponse
}

func (s ScheduleResponse) getDurations() ([]int) {
	return <-s.DurationResponse
}

func (s ScheduleResponse) getErrors() ([]error) {
	return <-s.ErrorResponse
}
