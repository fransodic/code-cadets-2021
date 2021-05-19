package http

import (
	"code-cadets-2021/homework_2/internal/domain/models"
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const axilisFeedCSVURL = "http://18.193.121.232/axilis-feed-2"

type AxilisOfferFeedCSV struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeedCSV(httpClient http.Client) *AxilisOfferFeedCSV {
	return &AxilisOfferFeedCSV{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeedCSV) Start(ctx context.Context) error {
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
		case <-time.After(time.Second):
			for _, odd := range content {
				a.writeToUpdatesChannel(odd)
			}
		}
	}
}

func decodeAndParseContent(content string) ([]models.Odd, error) {
	var parsedContent []models.Odd

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		values := strings.Split(line, ",")

		coeff, err := strconv.ParseFloat(values[3], 64)
		if err != nil {
			return []models.Odd{}, errors.WithMessage(err, "error parsing coefficient")
		}

		parsedContent = append(parsedContent, models.Odd{
			Id:          values[0],
			Name:        values[1],
			Match:       values[2],
			Coefficient: coeff,
			Timestamp:   time.Now(),
		})
	}

	return parsedContent, nil
}

func (a *AxilisOfferFeedCSV) getOddsFromServer() ([]models.Odd, error) {

	res, err := a.httpClient.Get(axilisFeedCSVURL)
	if err != nil {
		return []models.Odd{}, err
	}

	bodyContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []models.Odd{}, err
	}

	parsedContent, err := decodeAndParseContent(string(bodyContent))
	if err != nil {
		return []models.Odd{}, err
	}

	return parsedContent, nil
}

func (a *AxilisOfferFeedCSV) writeToUpdatesChannel(odd models.Odd) {
	a.updates <- odd
}

func (a *AxilisOfferFeedCSV) String() string {
	return "axilis offer feed 2"
}

func (a *AxilisOfferFeedCSV) GetUpdates() chan models.Odd {
	return a.updates
}
