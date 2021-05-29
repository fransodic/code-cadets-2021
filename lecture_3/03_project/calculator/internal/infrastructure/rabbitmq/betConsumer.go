package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// BetConsumer consumes bets from the desired RabbitMQ queue.
type BetConsumer struct {
	channel Channel
	config  ConsumerConfig
}

// NewBetConsumer creates and returns a new BetConsumer.
func NewBetConsumer(channel Channel, config ConsumerConfig) (*BetConsumer, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet consumer initialization failed")
	}

	return &BetConsumer{
		channel: channel,
		config:  config,
	}, nil
}

// Consume consumes messages until the context is cancelled. An error will be returned if consuming
// is not possible.
func (c *BetConsumer) Consume(ctx context.Context) (<-chan models.Bet, error) {
	msgs, err := c.channel.Consume(
		c.config.Queue,
		c.config.ConsumerName,
		c.config.AutoAck,
		c.config.Exclusive,
		c.config.NoLocal,
		c.config.NoWait,
		c.config.Args,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet consumer failed to consume messages")
	}

	bets := make(chan models.Bet)
	go func() {
		defer close(bets)
		for msg := range msgs {
			var bet models.Bet
			err := json.Unmarshal(msg.Body, &bet)
			if err != nil {
				log.Println("Failed to unmarshal bet message", msg.Body)
			}

			// Once context is cancelled, stop consuming messages.
			select {
			case bets <- bet:
			case <-ctx.Done():
				return
			}
		}
	}()

	return bets, nil
}
