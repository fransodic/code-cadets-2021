package handler

import (
	"context"
	"log"

	domainmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Handler handles bets and event updates.
type Handler struct {
	betRepository BetRepository
}

// New creates and returns a new Handler.
func New(betRepository BetRepository) *Handler {
	return &Handler{
		betRepository: betRepository,
	}
}

// HandleBets handles bets.
func (h *Handler) HandleBets(
	ctx context.Context,
	bets <-chan rabbitmqmodels.Bet,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for bet := range bets {
			log.Println("Processing bet, betId:", bet.Id)

			// Calculate the domain bet based on the incoming bet received.
			domainBet := domainmodels.BetCalculated{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			// Insert the domain bet into the repository.
			err := h.betRepository.InsertBet(ctx, domainBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}

		}
	}()

	return resultingBets
}

// HandleEventUpdates handles event updates.
func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventUpdate,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for eventUpdate := range eventUpdates {
			log.Println("Processing event update, betId:", eventUpdate.Id)

			// Fetch the domain bet.
			domainBets, exists, err := h.betRepository.GetBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch a bet which relates to this event update, error: ", err)
				continue
			}
			if !exists {
				log.Println("A bet which should be calculated does not exist, betId: ", eventUpdate.Id)
				continue
			}

			// Calculate payout based on event update.

			for _, domainBet := range domainBets {
				var resultingBet rabbitmqmodels.BetCalculated
				if eventUpdate.Outcome == "won" {
					resultingBet = rabbitmqmodels.BetCalculated{
						Id:     domainBet.Id,
						Status: eventUpdate.Outcome,
						Payout: domainBet.SelectionCoefficient * domainBet.Payment,
					}
				} else {
					resultingBet = rabbitmqmodels.BetCalculated{
						Id:     domainBet.Id,
						Status: eventUpdate.Outcome,
						Payout: 0 * domainBet.Payment,
					}
				}

				select {
				case resultingBets <- resultingBet:
				case <-ctx.Done():
					return
				}
			}

		}
	}()

	return resultingBets
}
