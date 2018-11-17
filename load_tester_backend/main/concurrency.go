package main

import (
	"fmt"
	"load_tester/requests"
	"net/http"
	"time"
)

func main() {
	//start := time.Now()
	requestCount := 100
	scheduleCount := 1
	maxConcurrent := 3
	client := &http.Client{}

	request := requests.NewGetRequest(
		"https://st-ccsd.accruenttest.net/MobileWebServices/apis/360facility/v1/workorders",
		map[string]string{"Authorization": "Bearer SXb8Jh8rfFIY-heK4eWRZtN5SQ5LQE3OP156yAJXZ9stjXyd7JWYHBp0XHZR--JHzfOIsDxQHWqpszOIG9GN2JfMXu23Z9f61wrneU0aNm6uRtux_Fx5BIr0JuqjVEaQduoqFUU4G-vhDBkr7e8tl-4bV0Q9Y0F2FYRD0sZ0rrdDevGyAwZ7p-0XtUFPhSlEbYadbblz9NHMp4z14uKbMBV61Baw4uf5kA25LKjSpHxxN8WJkOBLUTiRLOcASFbf15e7KGSg06m4wxkvoIK9Xebt6RFxd93oVta91w46ppD8V4-nWnui56ub1G8ZlJAGUmuM-kE4HGh3lLCswKJn_-rj99_CTTKK_iKeB8648zMaQa_gDFvbBL5QjvELx2AHZhcuA8nf4WmLw00nFfEAQ6dlkCA-jG4jnfJD7MCGMd9q-sGDW_oT3vkK4etp3gEX76U5SYxEnYremF0dQ2nHH4cMvGKYHovuJ8XSm4zeokG7Th5E1SZuTRW0J4pRiLcSJ9manNlKlEVU4tjQKqD7wuzkoB6R-gnlFHC56b6syC3uQlF7r-pCdAuyyy1mwsEE71X-fQ"},
		map[string]string{"$top": "100"},
		client,
	)


	//request := requests.NewGetRequest(
	//	"http://google.com",
	//	nil,
	//	nil,
	//	client,
	//)

	//request1 := requests.NewGetRequest(
	//	"http://facebook.com",
	//	nil,
	//	nil,
	//	client,
	//)
	schedule := requests.NewSchedule(time.Millisecond*10, requestCount, request, maxConcurrent)
	//schedule2 := requests.NewSchedule(time.Millisecond*300, requestCount, request1, maxConcurrent)
	resp := schedule.Run()
	//resp1 := schedule2.Run()

	responses := []*requests.ScheduleResponse{resp}
	for i := 0; i < scheduleCount; i++ {
		_, durations, _ := drainScheduleResponse(responses[i])
		//println("messages")
		//for idx, message := range messages {
		//	println(message)
		//	println(idx)
		//}
		total := 0
		for idx, dur := range durations {
			println(fmt.Sprintf("Duration  %d ms", dur / 1000000))
			total += dur
			println(idx)
		}

		print(fmt.Sprintf("Average duration : %d", (total / requestCount) / 1000000))
		//println("errors")
		//for idx, err := range errors {
		//	if err != nil {
		//		println(err.Error())
		//	}
		//	println(idx)
		//}
	}

}

func drainScheduleResponse(resp *requests.ScheduleResponse) ([]string, []int, []error) {
	msg := make(chan []string)
	dur := make(chan []int)
	errs := make(chan []error)
	go func() {
		msg <- <-resp.MessageResponse
		dur <- <-resp.DurationResponse
		errs <- <-resp.ErrorResponse
	}()

	return <-msg, <-dur, <-errs
}
