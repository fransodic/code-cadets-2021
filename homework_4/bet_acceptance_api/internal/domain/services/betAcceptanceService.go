package services

// BetAcceptanceService implements event related functions.
type BetAcceptanceService struct {
	betReceivedPublisher BetReceivedPublisher
}

// NewBetAcceptanceService creates a new instance of BetAcceptanceService.
func NewBetAcceptanceService(betReceivedPublisher BetReceivedPublisher) *BetAcceptanceService {
	return &BetAcceptanceService{
		betReceivedPublisher: betReceivedPublisher,
	}
}

// PublishBet sends event update message to the queues.
func (e BetAcceptanceService) PublishBet(customerId, selectionId string, selectionCoefficient, payment float64) error {
	return e.betReceivedPublisher.Publish(customerId, selectionId, selectionCoefficient, payment)
}
