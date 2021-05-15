package http

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(
	httpClient http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	// repeatedly:
	// - get odds from HTTP server
	// - write them to updates channel
	// - if context is finished, exit and close updates channel
	// (test your program from cmd/main.go)

	for {
		content, err := a.getOddsFromServer()
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			close(a.updates)
			return nil
		case <-time.After(time.Second * 3):
			for _, odd := range content {
				a.writeToUpdatesChannel(odd)
			}
		}
	}
}

func (a *AxilisOfferFeed) getOddsFromServer() ([]axilisOfferOdd, error) {

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

func (a *AxilisOfferFeed) writeToUpdatesChannel(odd axilisOfferOdd) {
	a.updates <- models.Odd{
		Id:          odd.Id,
		Name:        odd.Name,
		Match:       odd.Name,
		Coefficient: odd.Details.Price,
		Timestamp:   time.Time{},
	}
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
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
