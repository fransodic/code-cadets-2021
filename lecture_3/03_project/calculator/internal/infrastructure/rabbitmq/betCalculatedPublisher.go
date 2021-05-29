package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// BetCalculatedPublisher publishes calculated bets into the desired RabbitMQ queue.
type BetCalculatedPublisher struct {
	channel Channel
	config  PublisherConfig
}

// NewBetCalculatedPublisher creates and returns a new BetCalculatedPublisher.
func NewBetCalculatedPublisher(channel Channel, config PublisherConfig) (*BetCalculatedPublisher, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "calculated bet publisher initialization failed")
	}

	return &BetCalculatedPublisher{
		channel: channel,
		config:  config,
	}, nil
}

// Publish publishes messages until the context is cancelled.
func (p *BetCalculatedPublisher) Publish(ctx context.Context, betsCalculated <-chan models.BetCalculated) {
	go func() {
		for betCalculated := range betsCalculated {
			select {
			case <-ctx.Done():
				return
			default:
				betCalculatedJson, err := json.Marshal(&betCalculated)
				if err != nil {
					log.Println("Failed to marshal the following calculated bet:", betCalculated)
					continue
				}
				err = p.channel.Publish(
					p.config.Exchange,
					p.config.Queue,
					p.config.Mandatory,
					p.config.Immediate,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        betCalculatedJson,
					},
				)
				if err != nil {
					log.Println("Failed to publish the following calculated bet:", betCalculatedJson)
				}
			}
		}
	}()
}
