package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/api/option"
	"load_tester/load_tester_backend/requests"
	"net/http"
	"time"
)

type Server struct {
	httpClient *http.Client
	fireApp    *firebase.App
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
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.POST("schedule/request", s.CreateScheduleRequest)
	e.Start(":9000")
}

const MAX_CONCURRENT = 10

func (s Server) CreateScheduleRequest(c echo.Context) error {
	// todo validate input
	ctx := context.Background()
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

	// Done channel will send when all responses have been written
	// response channel is the channel all the responses will be written to
	respChannel, _ := schedule.Run()

	// firebase schedule record
	uid, _ := uuid.NewUUID()
	schedId := uid.String()
	fbSchedule := requests.Schedule{
		Id:           schedId,
		Name:         schedId,
		StartTime:    time.Now(),
		RequestCount: schedRequest.RequestCount,
	}

	_, err = s.createSchedule(fbSchedule)
	if err != nil {
		panic(err)
	}
	go func() {
		for resp := range respChannel {
			ref := s.getRequestCollection()
			docRef := ref.NewDoc()
			// important setting the key on this doc
			resp.Id = docRef.ID
			resp.ScheduleId = schedId
			res, err := docRef.Create(ctx, resp)
			if err != nil {
				panic(err)
			}
			print(res)
		}
	}()
	//<-doneChannel
	response := &requests.CreateScheduleResponse{
		ScheduleId: schedId,
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

func (s Server) createSchedule(schedule requests.Schedule) (*firestore.WriteResult, error) {
	collRef := s.getScheduleCollection()
	docRef := collRef.NewDoc()
	return docRef.Create(context.Background(), schedule)
}

func (s Server) getScheduleCollection() *firestore.CollectionRef {
	return s.getCollection("schedules")
}

func (s Server) getRequestCollection() *firestore.CollectionRef {
	return s.getCollection("requests")
}
