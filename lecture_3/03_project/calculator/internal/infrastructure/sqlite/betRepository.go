package sqlite

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	domainmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
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

func (r *BetRepository) queryInsertBet(ctx context.Context, bet storagemodels.BetCalculated) error {
	// always insert

	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateBet updates the provided bet in the database. An error is returned if the operation
// has failed.
func (r *BetRepository) UpdateBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.betMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *BetRepository) queryUpdateBet(ctx context.Context, bet storagemodels.BetCalculated) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
}

func (r *BetRepository) GetByID(ctx context.Context, id string) (storagemodels.BetCalculated, bool, error) {
	storageBet, err := r.queryGetBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return storagemodels.BetCalculated{}, false, nil
	}
	if err != nil {
		return storagemodels.BetCalculated{}, false, errors.Wrap(err, "bet repository failed to get bet with id "+id)
	}

	return storageBet, true, nil
}

func (r *BetRepository) queryGetBetByID(ctx context.Context, id string) (storagemodels.BetCalculated, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.BetCalculated{}, err
	}
	defer row.Close()

	// next and only result (id is unique)
	row.Next()

	var selectionId string
	var selectionCoefficient int
	var payment int

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
	if err != nil {
		return storagemodels.BetCalculated{}, err
	}

	return storagemodels.BetCalculated{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}

// GetBySelectionID fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetRepository) GetBySelectionID(ctx context.Context, id string) ([]domainmodels.Bet, error) {
	storageBets, err := r.queryGetBetsBySelectionID(ctx, id)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, errors.Wrap(err, "bet repository failed to get bets with selection id "+id)
	}

	var domainBets []domainmodels.Bet
	for _, storageBet := range storageBets {
		domainBet := r.betMapper.MapStorageBetToDomainBet(storageBet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, nil
}

func (r *BetRepository) queryGetBetsBySelectionID(ctx context.Context, id string) ([]storagemodels.BetCalculated, error) {
	rows, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+id+"';")

	if err != nil {
		return []storagemodels.BetCalculated{}, err
	}
	defer rows.Close()

	var results []storagemodels.BetCalculated
	// This will move to the "next" result and iterate through all rows.
	for rows.Next() {
		var selectionId string
		var selectionCoefficient int
		var payment int

		err = rows.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return []storagemodels.BetCalculated{}, err
		}

		results = append(results, storagemodels.BetCalculated{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		})
	}

	return results, nil
}
