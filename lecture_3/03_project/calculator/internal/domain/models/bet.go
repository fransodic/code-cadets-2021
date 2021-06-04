package models

// Bet represents a domain model representation of a calculated bet.
type Bet struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
}
