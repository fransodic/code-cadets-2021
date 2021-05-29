package services

import (
	"context"

	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

// BetService implements bet related functions
type BetService struct {
	betRepository BetRepository
}

func NewBetService(betRepository BetRepository) *BetService {
	return &BetService{
		betRepository: betRepository,
	}
}

func (b *BetService) GetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error) {
	return b.betRepository.GetBetByID(ctx, id)
}

func (b *BetService) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.Bet, error) {
	return b.betRepository.GetBetsByStatus(ctx, status)
}

func (b *BetService) GetBetsByCustomerID(ctx context.Context, customerId string) ([]domainmodels.Bet, error) {
	return b.betRepository.GetBetsByCustomerID(ctx, customerId)
}
