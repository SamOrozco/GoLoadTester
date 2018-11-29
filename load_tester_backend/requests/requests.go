package requests

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

type Request struct {
	RequestUrl  string
	StartTime   time.Time
	RequestType string
	Headers     map[string]string
	QueryParams map[string]string
	Client      *http.Client
}

type RequestResponse struct {
	Id          string `json:"id" firestore:"id"`
	ScheduleId  string `json:"schedule_id" firestore:"schedule_id"`
	ErrString   string `json:"err_string" firestore:"err_string"`
	Duration    int    `json:"duration" firestore:"duration"`
	RequestUrl  string `json:"request_url" firestore:"request_url"`
	RequestType string `json:"request_type" firestore:"request_type"`
	Message     string `json:"message" firestore:"message"`
}
type ScheduleRequest struct {
	RequestUrl    string            `json:"url"`
	RequestType   string            `json:"requestType"`
	Headers       map[string]string `json:"headers"`
	QueryParams   map[string]string `json:"queryParams"`
	RequestCount  int               `json:"requestCount"`
	IntervalCount int               `json:"intervalCount"`
	IntervalType  string            `json:"intervalType"`
	Block         bool              `json:"block"`
}

type CreateScheduleResponse struct {
	ScheduleId string `json:"schedule_id"`
}

func NewGetRequest(
	requestUrl string,
	headers map[string]string,
	params map[string]string,
	client *http.Client) *Request {
	return &Request{
		RequestUrl:  requestUrl,
		Headers:     headers,
		QueryParams: params,
		Client:      client,
	}
}

func (r Request) Run() (string, int, error) {
	r.StartTime = time.Now()

	// building request
	req, err := http.NewRequest(GET, r.RequestUrl, nil)
	if err != nil {
		return "", 0, err
	}

	if r.Headers != nil && len(r.Headers) > 0 {
		// set headers
		for key, value := range r.Headers {
			req.Header.Add(key, value)
		}
	}

	if r.QueryParams != nil && len(r.QueryParams) > 0 {
		queries := req.URL.Query()
		for key, value := range r.QueryParams {
			queries.Add(key, value)
		}
		req.URL.RawQuery = queries.Encode()
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != 200 {
		return "", 0, errors.New(fmt.Sprintf("Status code not 200, status code: %s", resp.StatusCode))
	}

	duration := int(time.Since(r.StartTime).Nanoseconds())
	bodyString, err := readBody(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return bodyString, duration, nil
}

func readBody(bodyReader io.ReadCloser) (string, error) {
	bytes, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
