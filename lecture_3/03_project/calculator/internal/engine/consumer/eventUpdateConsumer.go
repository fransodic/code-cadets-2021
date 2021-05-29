package consumer

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type EventUpdateConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error)
}
