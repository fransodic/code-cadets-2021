package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "code-cadets-2021/homework_4/bets_api/internal/domain/models"
	storagemodels "code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on cals_bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// GetBetByID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	if !row.Next() {
		return storagemodels.Bet{}, sql.ErrNoRows
	}

	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payout int

	err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payout)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	return storagemodels.Bet{
		Id:                   id,
		CustomerId:           customerId,
		Status:               status,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
		Payout:               payout,
	}, nil
}

// GetBetsByStatus fetches all bets from the database that have provided status
// and returns them. If no bets exist exist, an error will not be returned.
func (r *BetRepository) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.Bet, error) {
	storageBets, err := r.queryGetBetsByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, errors.Wrap(err, "bets repository failed to get a bet with status "+status)
	}

	var domainBets []domainmodels.Bet
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, nil
}

func (r *BetRepository) queryGetBetsByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE status='"+status+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	// A loop over all returned rows.
	for row.Next() {
		var id string
		var customerId string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payout int

		err = row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payout)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		bets = append(bets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               status,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return bets, nil
}

// GetBetsByCustomerID fetches all bets from the database that have provided status
// and returns them. If no bets exist, an error will not be returned.
func (r *BetRepository) GetBetsByCustomerID(ctx context.Context, customerId string) ([]domainmodels.Bet, error) {
	storageBets, err := r.queryGetBetsByCustomerID(ctx, customerId)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, errors.Wrap(err, "bets repository failed to get bets with customer ID "+customerId)
	}

	var domainBets []domainmodels.Bet
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, nil
}

func (r *BetRepository) queryGetBetsByCustomerID(ctx context.Context, customerId string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE customer_id='"+customerId+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	// A loop over all returned rows.
	for row.Next() {
		var id string
		var status string
		var selectionId string
		var selectionCoefficient int
		var payment int
		var payout int

		err = row.Scan(&id, &status, &customerId, &selectionId, &selectionCoefficient, &payment, &payout)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		bets = append(bets, storagemodels.Bet{
			Id:                   id,
			CustomerId:           customerId,
			Status:               customerId,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
			Payout:               payout,
		})
	}

	return bets, nil
}
