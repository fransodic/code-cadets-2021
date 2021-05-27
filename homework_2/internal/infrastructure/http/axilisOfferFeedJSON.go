package http

import (
	"code-cadets-2021/homework_2/internal/domain/models"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeedJSON struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeedJSON(httpClient http.Client) *AxilisOfferFeedJSON {
	return &AxilisOfferFeedJSON{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeedJSON) Start(ctx context.Context) error {
	// repeatedly:
	// - get odds from HTTP server
	// - write them to updates channel
	// - if context is finished, exit and close updates channel
	// (test your program from cmd/main.go)

	for {
		content, err := a.getOddsFromServer()
		if err != nil {
			//fmt.Println("http err")
			return err
		}

		select {
		case <-ctx.Done():
			close(a.updates)
			//fmt.Println("feed finish")
			return nil
		case <-time.After(time.Second):
			//fmt.Println("prosla sekunda")
			for _, odd := range content {
				//fmt.Println("pisem odd")
				a.writeToUpdatesChannel(odd)
			}
		}
	}
}

func (a *AxilisOfferFeedJSON) getOddsFromServer() ([]axilisOfferOdd, error) {

	res, err := a.httpClient.Get(axilisFeedURL)
	if err != nil {
		return []axilisOfferOdd{}, err
	}

	bodyContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []axilisOfferOdd{}, err
	}

	var decodedContent []axilisOfferOdd
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		return []axilisOfferOdd{}, err
	}

	return decodedContent, nil
}

func (a *AxilisOfferFeedJSON) writeToUpdatesChannel(odd axilisOfferOdd) {
	a.updates <- models.Odd{
		Id:          odd.Id,
		Name:        odd.Name,
		Match:       odd.Name,
		Coefficient: odd.Details.Price,
		Timestamp:   time.Now(),
	}
}

func (a *AxilisOfferFeedJSON) String() string {
	return "axilis offer feed"
}

func (a *AxilisOfferFeedJSON) GetUpdates() chan models.Odd {
	return a.updates
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
