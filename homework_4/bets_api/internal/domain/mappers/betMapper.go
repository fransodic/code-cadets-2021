package mappers

import (
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
	storagemodels "code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"
)

// BetMapper maps storage bets to domain bets and vice versa.
type BetMapper struct {
}

// NewBetMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
func (m *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet {
	return domainmodels.Bet{
		Id:                   storageBet.Id,
		Status:               storageBet.Status,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
		Payout:               float64(storageBet.Payout) / 100,
	}
}

func (m *BetMapper) MapResultToDto(bet domainmodels.Bet) models.BetResponseDto {
	return models.BetResponseDto{
		Id:                   bet.Id,
		Status:               bet.Status,
		SelectionId:          bet.SelectionId,
		SelectionCoefficient: bet.SelectionCoefficient,
		Payment:              bet.Payment,
		Payout:               bet.Payout,
	}
}
