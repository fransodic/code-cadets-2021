package sqlite

import (
	domainmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapDomainBetToStorageBet(domainBet domainmodels.BetCalculated) storagemodels.BetCalculated
	MapStorageBetToDomainBet(storageBet storagemodels.BetCalculated) domainmodels.BetCalculated
}
