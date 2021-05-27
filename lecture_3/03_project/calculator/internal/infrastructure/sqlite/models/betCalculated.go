package models

// BetCalculated is a storage model representation of a calculated bet.
type BetCalculated struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient int
	Payment              int
}
