package publisher

import (
	"context"

	rabbitmqmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Publisher offers methods for publishing into output queues.
type Publisher struct {
	betPublisher BetCalculatedPublisher
}

// New creates and returns a new Publisher.
func New(betPublisher BetCalculatedPublisher) *Publisher {
	return &Publisher{
		betPublisher: betPublisher,
	}
}

// PublishBets publishes into bets queue.
func (p *Publisher) PublishBetsCalculated(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated) {
	p.betPublisher.Publish(ctx, bets)
}
