package engine

import (
	"context"

	rabbitmqmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher interface {
	PublishBetsCalculated(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated)
}
