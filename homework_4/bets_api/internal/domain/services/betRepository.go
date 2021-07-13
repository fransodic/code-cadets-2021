package services

import (
	"context"

	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

type BetRepository interface {
	GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error)
	GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.Bet, error)
	GetBetsByCustomerID(ctx context.Context, customerId string) ([]domainmodels.Bet, error)
}
