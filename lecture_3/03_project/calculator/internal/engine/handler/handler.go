package handler

import (
	"context"
	"log"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
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

			// Fetch the domain bet.
			domainBet, exists, err := h.betRepository.GetBetByID(ctx, bet.Id)
			if err != nil {
				log.Println("Failed to fetch a bet, error: ", err)
				continue
			}

			// Insert the bet if it does not already exist.
			if !exists {
				// Calculate the bet based on the incoming bet.
				domainBet = domainmodels.Bet{
					Id:                   bet.Id,
					SelectionId:          bet.SelectionId,
					SelectionCoefficient: bet.SelectionCoefficient,
					Payment:              bet.Payment,
				}

				// Insert the bet into the repository.
				err := h.betRepository.InsertBet(ctx, domainBet)
				if err != nil {
					log.Println("Failed to insert bet, error: ", err)
					continue
				}
			} else {
				// Update the domain bet based on incoming changes.
				domainBet.SelectionId = bet.SelectionId
				domainBet.SelectionCoefficient = bet.SelectionCoefficient
				domainBet.Payment = bet.Payment

				// Update the domain bet into the repository.
				err = h.betRepository.UpdateBet(ctx, domainBet)
				if err != nil {
					log.Println("Failed to update bet, error: ", err)
					continue
				}
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
			log.Println("Processing event update, selectionId:", eventUpdate.Id)

			// Fetch all domain bets with this selectionID.
			domainBets, err := h.betRepository.GetBetsBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch bets, error: ", err)
				continue
			}
			if len(domainBets) == 0 {
				log.Println("There are no bets with selectionId: ", eventUpdate.Id)
				continue
			}

			for _, domainBet := range domainBets {
				var resultingBet rabbitmqmodels.BetCalculated

				// Calculate the resulting bet, which should be published.
				if eventUpdate.Outcome == "won" {
					resultingBet = rabbitmqmodels.BetCalculated{
						Id:     domainBet.Id,
						Status: eventUpdate.Outcome,
						Payout: domainBet.Payment * domainBet.SelectionCoefficient,
					}
				} else if eventUpdate.Outcome == "lost" {
					resultingBet = rabbitmqmodels.BetCalculated{
						Id:     domainBet.Id,
						Status: eventUpdate.Outcome,
						Payout: 0,
					}
				} else {
					log.Printf("Unknown outcome on the following event update: %s", eventUpdate)
				}

				log.Printf("%v\n", resultingBet)

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
