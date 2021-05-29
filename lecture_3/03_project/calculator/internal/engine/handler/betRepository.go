package handler

import (
	storagemodels "code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
	"context"

	domainmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type BetRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateBet(ctx context.Context, bet domainmodels.Bet) error
	GetBySelectionID(ctx context.Context, id string) ([]domainmodels.Bet, error)
	GetByID(ctx context.Context, id string) (storagemodels.BetCalculated, bool, error)
}
