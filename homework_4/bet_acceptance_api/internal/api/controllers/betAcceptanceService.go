package controllers

// BetAcceptanceService implements bet related functions.
type BetAcceptanceService interface {
	PublishBet(customerId, selectionId string, selectionCoefficient, payment float64) error
}
