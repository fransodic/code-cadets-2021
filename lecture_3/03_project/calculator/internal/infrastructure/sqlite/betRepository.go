package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
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

// InsertBet inserts the provided bet into the database. An error is returned if the operation
// has failed.
func (r *BetRepository) InsertBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryInsertBet(ctx context.Context, bet storagemodels.Bet) error {
	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateBet updates the provided bet in the database. An error is returned if the operation has failed.
func (r *BetRepository) UpdateBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryUpdateBet(ctx context.Context, bet storagemodels.Bet) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
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

	var selectionId string
	var selectionCoefficient int
	var payment int

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	return storagemodels.Bet{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}

// GetBetsBySelectionID fetches all bets from the database that have provided selectionId
// and returns them. The second returned value indicates whether the bet exists in DB.
// If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBetsBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.Bet, error) {
	storageBets, err := r.queryGetBetBySelectionID(ctx, selectionId)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, errors.Wrap(err, "bets repository failed to get a bet with selection id "+selectionId)
	}

	var domainBets []domainmodels.Bet
	for _, storageBet := range storageBets {
		domainBets = append(domainBets, r.betMapper.MapStorageBetToDomainBet(storageBet))
	}

	return domainBets, nil
}

func (r *BetRepository) queryGetBetBySelectionID(ctx context.Context, selectionId string) ([]storagemodels.Bet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+selectionId+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	// A loop over all returned rows.
	for row.Next() {
		var id string
		var selectionCoefficient int
		var payment int

		err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return []storagemodels.Bet{}, err
		}

		bets = append(bets, storagemodels.Bet{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		})
	}

	return bets, nil
}
