package controllers

// BetRequestValidator validates event update requests.
type BetRequestValidator interface {
	StatusIsValid(status string) bool
}
