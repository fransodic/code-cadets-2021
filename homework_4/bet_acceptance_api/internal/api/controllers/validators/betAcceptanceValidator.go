package validators

import "code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"

// BetAcceptanceValidator validates event update requests.
type BetAcceptanceValidator struct {
	coefficientUpperBound float64
	paymentLowerBound     float64
	paymentUpperBound     float64
}

// NewBetAcceptanceValidator creates a new instance of BetAcceptanceValidator.
func NewBetAcceptanceValidator(coeffUpperBound, paymentLowerBound, paymentUpperBound float64) *BetAcceptanceValidator {
	return &BetAcceptanceValidator{
		coefficientUpperBound: coeffUpperBound,
		paymentLowerBound:     paymentLowerBound,
		paymentUpperBound:     paymentUpperBound,
	}
}

// BetIsValid checks if received bet is valid.
// Id is non-default
// Coefficient is non-default and lower than given upper bound
// Payment is non-default and between given upper and lower bound
func (b *BetAcceptanceValidator) BetIsValid(betAcceptanceRequestDto models.BetAcceptanceRequestDto) bool {
	if b.isValidCustomerID(betAcceptanceRequestDto.CustomerId) && b.isValidPayment(betAcceptanceRequestDto.Payment) && b.isValidCoefficient(betAcceptanceRequestDto.SelectionCoefficient) {
		return true
	}

	return false
}

func (b *BetAcceptanceValidator) isValidCoefficient(coefficient float64) bool {
	return coefficient != float64(0) && coefficient <= b.coefficientUpperBound
}

func (b *BetAcceptanceValidator) isValidPayment(payment float64) bool {
	return payment != float64(0) && payment >= b.paymentLowerBound && payment <= b.paymentUpperBound
}

func (b *BetAcceptanceValidator) isValidCustomerID(id string) bool {
	return id != ""
}
