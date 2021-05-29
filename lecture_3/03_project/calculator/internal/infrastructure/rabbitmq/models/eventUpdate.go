package models

// EventUpdate represents a DTO for event updates.
type EventUpdate struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}
