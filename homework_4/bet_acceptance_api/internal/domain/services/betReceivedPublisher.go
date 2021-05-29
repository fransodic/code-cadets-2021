package services

// BetReceivedPublisher handles event update queue publishing.
type BetReceivedPublisher interface {
	Publish(customerId, selectionId string, selectionCoefficient, payment float64) error
}
