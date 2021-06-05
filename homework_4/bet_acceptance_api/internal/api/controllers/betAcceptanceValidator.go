package controllers

import "code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"

// BetAcceptanceValidator validates bet acceptance requests.
type BetAcceptanceValidator interface {
	BetIsValid(eventUpdateRequestDto models.BetAcceptanceRequestDto) bool
}
