package main

import (
	"concurrency_fun/requests"
	"fmt"
	"strings"
	"time"
)

func main() {
	requestCount := 5
	messageChannel := make(chan string)
	durationChannel := make(chan int64)
	errorChannel := make(chan error)
	for i := 0; i < requestCount; i++ {
		request := requests.NewGetRequest(
			"http://google.com",
			messageChannel,
			durationChannel,
			errorChannel)

		go request.Run()
	}

	// read errors from our error channel
	errorResponseChannel := make(chan string)
	go func() {
		var sb strings.Builder
		for reqError := range errorChannel {
			sb.WriteString(reqError.Error() + "\r\n")
		}
		errorResponseChannel <- sb.String()
	}()

	// read all messages from our messages channel
	messageBodyResponse := make(chan string)
	go func() {
		var sb strings.Builder
		for message := range messageChannel {
			sb.WriteString(message + "\r\n")
		}
		messageBodyReponse <- sb.String()
	}()

	total := int64(0)
	for i := 0; i < requestCount; i++ {
		total += <-durationChannel
	}

	// wait for concurrent error response
	select {
	case errorMsg := <-errorResponseChannel:
		println(errorMsg)
	case <-time.After(1 * time.Second):
		println("No errors")
	}

	println(<-messageBodyReponse)

	print(fmt.Sprintf("Average request time %d ms", (total/int64(requestCount))/1000000))
}
