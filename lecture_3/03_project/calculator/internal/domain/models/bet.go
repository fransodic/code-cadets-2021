package models

// BetCalculated represents a domain model representation of a calculated bet.
type BetCalculated struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
}
