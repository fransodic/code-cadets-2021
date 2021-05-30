package rabbitmq

import (
	"encoding/json"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"time"

	"code-cadets-2021/homework_4/bet_acceptance_api/internal/infrastructure/rabbitmq/models"
	"github.com/streadway/amqp"
)

const contentTypeTextPlain = "text/plain"

// BetReceivedPublisher handles event update queue publishing.
type BetReceivedPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewEventUpdatePublisher create a new instance of BetReceivedPublisher.
func NewEventUpdatePublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetReceivedPublisher {
	return &BetReceivedPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

// Publish publishes an event update message to the queue.
func (p *BetReceivedPublisher) Publish(customerId, selectionId string, selectionCoefficient, payment float64) error {
	rand.Seed(time.Now().UnixNano())

	id, err := getRandomUUID()
	if err != nil {
		return errors.Wrap(err, "failed to generate a random UUID")
	}

	betReceived := &models.BetReceivedDto{
		Id:                   id,
		CustomerId:           customerId,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}

	eventUpdateJson, err := json.Marshal(betReceived)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        eventUpdateJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", eventUpdateJson)
	return nil
}

func getRandomUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
