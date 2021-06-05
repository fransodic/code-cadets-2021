package controllers

import (
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
	storagemodels "code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet
	MapResultToDto(bet domainmodels.Bet) models.BetResponseDto
}
