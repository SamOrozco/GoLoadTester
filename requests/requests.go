package requests

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestType string

const (
	GET  RequestType = "GET"
	POST RequestType = "POST"
)

type Request struct {
	requestUrl      string
	startTime       time.Time
	messageChannel  chan string
	durationChannel chan int64
	errorChannel    chan error
	requestType     RequestType
}

func NewGetRequest(
	requestUrl string,
	messageChannel chan string,
	durationChannel chan int64,
	errorChannel chan error) *Request {
	return &Request{
		requestUrl:      requestUrl,
		messageChannel:  messageChannel,
		durationChannel: durationChannel,
		errorChannel:    errorChannel,
	}
}

func (r Request) Run() {
	r.startTime = time.Now()

	resp, err := http.Get(r.requestUrl)
	if err != nil {
		// todo handle
		r.errorChannel <- err
	}
	if resp.StatusCode != 200 {
		// todo handle
		r.errorChannel <- errors.New("Status code not 200")
	}

	r.durationChannel <- time.Since(r.startTime).Nanoseconds()
	// ready body and write body to message channel
	bodyString, err := readBody(resp.Body)
	if err != nil {
		r.errorChannel <- err
	}
	r.messageChannel <- bodyString
}

func readBody(bodyReader io.ReadCloser) (string, error) {
	bytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
