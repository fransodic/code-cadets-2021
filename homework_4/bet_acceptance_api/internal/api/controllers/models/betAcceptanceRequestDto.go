package models

// BetAcceptanceRequestDto Update request dto model.
type BetAcceptanceRequestDto struct {
	CustomerId           string  `json:"customerId"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
}
