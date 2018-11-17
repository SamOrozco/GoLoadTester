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
	requestUrl  string
	startTime   time.Time
	requestType string
	headers     map[string]string
	queryParams map[string]string
	client      *http.Client
}

func NewGetRequest(
	requestUrl string,
	headers map[string]string,
	params map[string]string,
	client *http.Client) *Request {
	return &Request{
		requestUrl:  requestUrl,
		headers:     headers,
		queryParams: params,
		client:      client,
	}
}

func (r Request) Run() (string, int, error) {
	println("Running")
	r.startTime = time.Now()

	// building request
	req, err := http.NewRequest(GET, r.requestUrl, nil)
	if err != nil {
		return "", 0, err
	}

	if r.headers != nil && len(r.headers) > 0 {
		// set headers
		for key, value := range r.headers {
			req.Header.Add(key, value)
		}
	}

	if r.queryParams != nil && len(r.queryParams) > 0 {
		queries := req.URL.Query()
		for key, value := range r.queryParams {
			queries.Add(key, value)
		}
		req.URL.RawQuery = queries.Encode()
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != 200 {
		return "", 0, errors.New(fmt.Sprintf("Status code not 200, status code: %s", resp.StatusCode))
	}

	duration := int(time.Since(r.startTime).Nanoseconds())
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
