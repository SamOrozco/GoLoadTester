package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	context2 "golang.org/x/net/context"
	"google.golang.org/api/option"
	"load_tester/load_tester_backend/pulse"
	"load_tester/load_tester_backend/requests"
	"math"
	"net/http"
	"sync/atomic"
	"time"
)

type Server struct {
	httpClient *http.Client
	fireApp    *firebase.App
	// this will be a channel that if written to once will cancel the running schedule
	// the schedule creates it itself `CreateScheduleRequest` when it starts and removes it when the schedule is done
	stopScheduleMap map[string]chan bool
	countBuffer     int
	errorBuffer     int
}

const serviceAccountJson = `{"type":"service_account","project_id":"load-tester-orozco","private_key_id":"daf9558be8f3bfea5b6f74ca745cfb9012d607a5","private_key":"-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDOYN8xenXgyXuD\nXp/W/2lA9HI67+d8EHi43eG+rzaLyIfRZIQwMld2GU87w4cenZlMBbblQnNlX7aa\nu9skOXbZRpqfhqPpDdGL2ocDKEjiDpYhP06WCb+d7Zb02E3tCvQEpqUOMekpg/3u\n5oXx4JHRzzHFE+MRPs3c9LrTJ7WJGM0a8RYb07hzOEu2coE8By8VsKynQ0I4f9tA\ny75MTnZeEa/cTh6D+AfSvfp2sorsrcQ74EuITD04RVyO3T+BBPY/CaOInA3PaszN\n91BPPVu5y4KT5LWOJMEoY3n2DbXLswgD+mUtVPDhL/7wPKnzq4wrxWUcC928d0DH\nZqgZThgtAgMBAAECggEAE6TrPybvHH/2hp6AVsiGCMhGgpn+WJ6qwby2YF5bh0+w\nWeaYTIwmK216T2/WUebiG3c5yGBxt9cLUcw1g6EH++OOvpuOyOq8pPQeliwvT9E1\nSiHgBZkU87U/VrwAzYzq2vHP0+Y9UZXTmyS3NZBQ4P3jY2UitSwGEBT9dDGbQ9rR\nkWSdcsw6AIXJ7NpeTtBFF5Sg+5KqpAxgc/rtBtIfT1OGcA88ST1usZaH5bzakzDI\nXVaP+cEIwRvOjQPjMJgBiHqTlFDrMcWE4ljufdB1SPzbPgiXSGGkRnwyzaZVV7dQ\nT8vpCvEWX8ME4930WKi6dNE/XOIyTZHlSbdhcUar8wKBgQDo8szBE9gJ6ItY8gHH\nRh+WL6kEQOVJubFsH9mUKu0wd8ZbgevWn4+fhVHyExi+/rsgse0d5SKGOI6yT0im\nG2K2rHzRBEJBsZ6t5qvI+IOckPzWr/umhal+/eHXlQxGX9nyfVuUz64NnQGGqhVy\nyoV9GE8KrP2SpPL9aetIrIEsiwKBgQDizPub1qYNGlRs4EP5nAInC6xFZt344Fou\nW0d3zfw53S1Ocm6F8vXd+ik110oWq8YFkO3n3fu2E5ZCXKT/wKy6buCCbfN7qvjR\nbEdGMex1WTZZwEFmYExk0TIcpm+Pajv4WbIL86jtOyMYnO1Lsgifj6CPWndV2LNH\nFNFY2H/NJwKBgCfS+nY3skNfYhM+rFcpAdVnKHn1rj1AZbiaaVQPUqVqlJqoR21V\nsfQPksbTxPNCETOxt1vZBGH14U9ShpT/MY0RR/VjyP/6IaqROOUOt1FY6CC7zUXW\nq/dt5+I3NPYDBPuTushiCNuC7/PZt/j77n5IXZrxoBgl4vS32uV9uEepAoGBAN/m\nk4KenqGS28P8hr7WKD3ZxUFis0JuYjkffeelYBrT4lVunPP6DoYM4EA6APdurvhx\nwpxERqnSnmV3RqEB2sPWkCfTWis9d5Rv+9Etmg2jfAeQyD/EU8a3y2wDV2FS9E9S\nZXZtHtjMp/I1ggJXiTHUviKnoeYLH7nNsHqwiULDAoGBAKpD/8qgPe0dglnYm6vK\nSqsSrkSi4dW/2Bh9XC0+np0yDSg/IlHCC53KZupTVI2JKuf48lzwTL4Lj7e6JB21\nuSAwFYLQPhSkQDBX/WmqLuHVMHMYFhqd6LHU2pvId+6Xv9jVDWT9uVfV9Kxu60UQ\nV868m36HAxyHuMow0GJ1+yRo\n-----END PRIVATE KEY-----\n","client_email":"firebase-adminsdk-zg0yt@load-tester-orozco.iam.gserviceaccount.com","client_id":"111349850986428698651","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_x509_cert_url":"https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-zg0yt%40load-tester-orozco.iam.gserviceaccount.com"}`

func main() {
	// CREATE FIREBASE APP
	opt := option.WithCredentialsJSON([]byte(serviceAccountJson));
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}
	server := Server{httpClient: &http.Client{}, fireApp: app}
	server.Run()
}

func (s Server) Run() {
	// init map
	s.stopScheduleMap = make(map[string]chan bool)
	// how many counts will increment before we publish
	// we will publish count every 5 counts
	s.countBuffer = 5
	// same as count but for errors
	s.errorBuffer = 1
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://connector2.ngrok.io"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.POST("schedule/request", s.CreateScheduleRequest)
	e.GET("schedule/:schedule_id/stop", s.StopRunningSchedule)
	e.Start(":9000")
}

const MAX_CONCURRENT = 10

// This method will take a schedule id as a path variable. It will then check to see if this schedule is running
// if it is running stop the schedule, if not, return not found
func (s Server) StopRunningSchedule(c echo.Context) error {
	schedId := c.Param("schedule_id")
	// if we found a running entry in the map lets
	// write a term to it
	if term, ok := s.stopScheduleMap[schedId]; ok {
		term <- true
	}
	return c.NoContent(200)
}

// This method takes a create schedule request starts the request
// then returns a ok response if everything was `started` correctly, not run.
// it will then write it's results to firebase
func (s Server) CreateScheduleRequest(c echo.Context) error {
	// todo validate inputs
	context := context2.Background()
	schedRequest := &requests.ScheduleRequest{}
	err := c.Bind(schedRequest)
	if err != nil {
		return err
	}
	// building our request
	request := requests.NewGetRequest(schedRequest.RequestUrl,
		schedRequest.Headers,
		schedRequest.QueryParams,
		s.httpClient)

	// we need to convert out string intervalType and interval count to
	// a go Duration
	requestInterval, err := s.getDurationFromRequest(*schedRequest)
	if err != nil {
		return err
	}
	// building our schedule
	schedule := requests.NewSchedule(
		requestInterval,
		schedRequest.RequestCount,
		request,
		MAX_CONCURRENT)
	// fire base schedule record
	uid, _ := uuid.NewUUID()
	// we have our schedule uuid
	schedId := uid.String()
	s.stopScheduleMap[schedId] = make(chan bool)
	// Done channel will send when all responses have been written
	// response channel is the channel all the responses will be written to
	respChannel, doneChannel := schedule.Run(s.stopScheduleMap[schedId])
	donePulse := pulse.NewPulse(doneChannel)

	// if name is null we are just going to set to the schedId
	schedName := schedRequest.Name
	if len(schedName) < 1 {
		schedName = schedId
	}

	fbSchedule := requests.Schedule{
		Id:           schedId,
		Name:         schedName,
		StartTime:    time.Now().Unix(),
		RequestCount: schedRequest.RequestCount,
	}

	refId, err := s.createSchedule(fbSchedule)
	if err != nil {
		panic(err)
	}

	// reading responses in asynchronously.
	// this go routine will end when <- doneChannel has been called
	doneSignal := donePulse.Subscribe()

	// REPORTING VARIABLES
	totalDuration := int32(0)
	sentRequestCount := int32(0)
	longestDuration := int32(0)
	shortestDuration := int32(math.MaxInt32)
	go func() {
		for {
			select {
			case resp := <-respChannel:
				s.createRequestResponse(resp, refId)
				// add to total
				atomic.AddInt32(&totalDuration, int32(resp.Duration))
				// add to sent
				atomic.AddInt32(&sentRequestCount, int32(1))
				// update sent
				s.updateSentRequestCount(s.getSchedule(refId), sentRequestCount)
				// update average
				s.updateAverageDuration(s.getSchedule(refId), totalDuration/sentRequestCount)
				// new longest dur
				if resp.Duration > int(longestDuration) { // new longest dur
					atomic.StoreInt32(&longestDuration, int32(resp.Duration))
					// update firebase
					s.updateLongestDuration(s.getSchedule(refId), longestDuration)
				}
				// new shortest dur
				if resp.Duration < int(shortestDuration) {
					atomic.StoreInt32(&shortestDuration, int32(resp.Duration))
					// update firebase
					s.updateShortestDuration(s.getSchedule(refId), shortestDuration)
				}
			case _ = <-doneSignal:
				return
			}
		}
	}()

	// this will block until the done chan channel is written to
	waitOnSchedule := make(chan bool)
	doneSignal1 := donePulse.Subscribe()
	go func() {
		// THIS IS WHERE WE CAN DO THINGS ON SCHEDULE COMPLETION
		<-doneSignal1
		waitOnSchedule <- true

		// UPDATING SCHEDULE WITH FINAL DATA
		scheduleDoc := s.getSchedule(refId)
		_, err := scheduleDoc.Set(context, map[string]int64{"end_time": time.Now().Unix()}, firestore.MergeAll)
		if err != nil {
			panic(err)
		}
	}()

	// If the user sent `block = true` in their request we are
	// going to wait here else we return an async response and keep working
	if schedRequest.Block {
		<-waitOnSchedule
	} else {
		go func() {
			<-waitOnSchedule
		}()
	}
	response := &requests.CreateScheduleResponse{
		ScheduleId: refId,
	}
	return c.JSON(http.StatusOK, response)
}

func (s Server) getDurationFromRequest(req requests.ScheduleRequest) (time.Duration, error) {
	timeType := req.IntervalType
	intervalCount := req.IntervalCount
	switch timeType {
	case "Milliseconds":
		return time.Duration(int64(time.Millisecond) * int64(intervalCount)), nil
	case "Seconds":
		return time.Duration(int64(time.Second) * int64(intervalCount)), nil
	case "Minutes":
		return time.Duration(int64(time.Minute) * int64(intervalCount)), nil
	default:
		return time.Second * 30, nil
	}
}

func (s Server) getCollection(path string) (*firestore.CollectionRef) {
	store, err := s.fireApp.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	return store.Collection(path)
}

func (s Server) createSchedule(schedule requests.Schedule) (string, error) {
	collRef := s.getScheduleCollection()
	docRef := collRef.NewDoc()
	_, err := docRef.Create(context.Background(), schedule)
	if err != nil {
		return "", err
	}
	_, err = docRef.Set(context.Background(), map[string]string{"id": docRef.ID}, firestore.MergeAll)
	if err != nil {
		return "", err
	}
	return docRef.ID, err
}

func (s Server) getScheduleCollection() *firestore.CollectionRef {
	return s.getCollection("schedules")
}

func (s Server) getSchedule(scheduleId string) (*firestore.DocumentRef) {
	return s.getScheduleCollection().Doc(scheduleId)
}

func (s Server) getRequestCollection() *firestore.CollectionRef {
	return s.getCollection("requests")
}

// we're going to go simple for now
func (s Server) handleRequestResponse(resp requests.RequestResponse, scheduleId string) error {
	// assume if error string is not empty we got an errors
	return s.createRequestResponse(resp, scheduleId)
}

func (s Server) createRequestResponse(resp requests.RequestResponse, scheduleId string) error {
	ctx := context.Background()
	ref := s.getRequestCollection()
	docRef := ref.NewDoc()
	// important setting the key on this doc
	resp.Id = docRef.ID
	resp.ScheduleId = scheduleId
	_, err := docRef.Create(ctx, resp)
	return err
}

func (s Server) updateSentRequestCount(docRef *firestore.DocumentRef, count int32) {
	_, err := docRef.Set(context.Background(), map[string]int32{"current_request_count": count}, firestore.MergeAll)
	if err != nil {
		panic(err)
	}
}

func (s Server) updateAverageDuration(docRef *firestore.DocumentRef, duration int32) {
	_, err := docRef.Set(context.Background(), map[string]int32{"average_duration": duration}, firestore.MergeAll)
	if err != nil {
		panic(err)
	}
}

func (s Server) updateLongestDuration(docRef *firestore.DocumentRef, duration int32) {
	_, err := docRef.Set(context.Background(), map[string]int32{"longest_duration": duration}, firestore.MergeAll)
	if err != nil {
		panic(err)
	}
}


func (s Server) updateShortestDuration(docRef *firestore.DocumentRef, duration int32) {
	_, err := docRef.Set(context.Background(), map[string]int32{"shortest_duration": duration}, firestore.MergeAll)
	if err != nil {
		panic(err)
	}
}