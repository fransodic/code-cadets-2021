package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/rand"
	stdhttp "net/http"
	"time"
)

const eventUpdateEndpoint = "http://127.0.0.1:8080/event/update"
const betsAPI = "http://127.0.0.1:9090/bets"
const activeBetsParam = "?status=active"

type Bet struct {
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
	Payout               float64 `json:"payout"`
}

func main() {
	httpClient := stdhttp.Client{Timeout: time.Second * 10}

	responseData, err := getActiveOdds(httpClient)
	if err != nil {
		log.Fatalln("error getting data: ", err)
	}

	events := getDistinctEvents(responseData)

	eventUpdates := generateRandomOutcomes(events)

	for _, eventUpdate := range eventUpdates {
		requestBody, err := json.Marshal(eventUpdate)
		post, err := httpClient.Post(eventUpdateEndpoint, "text/plain", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatalln("error sending event updates: ", err)
		}
		post.Body.Close()
	}
}

func generateRandomOutcomes(events map[string]struct{}) []EventUpdateDto {
	outcomes := [2]string{"won", "lost"}
	rand.Seed(time.Now().Unix())

	var eventUpdates []EventUpdateDto
	for event, _ := range events {
		eventUpdates = append(eventUpdates, EventUpdateDto{
			id:      event,
			outcome: outcomes[rand.Intn(len(outcomes))],
		})
	}

	return eventUpdates
}

func getDistinctEvents(data []Bet) map[string]struct{} {
	var events map[string]struct{}
	for _, bet := range data {
		events[bet.SelectionId] = struct{}{}
	}
	return events
}

func getActiveOdds(httpClient stdhttp.Client) ([]Bet, error) {
	res, err := httpClient.Get(betsAPI + activeBetsParam)
	if err != nil {
		return nil, errors.Wrap(err, "error getting data from bets API")
	}

	bodyContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}

	var decodedContent []Bet
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		return []Bet{}, errors.Wrap(err, "error unmarshalling response body")
	}

	return decodedContent, nil
}

type EventUpdateDto struct {
	id      string
	outcome string
}
