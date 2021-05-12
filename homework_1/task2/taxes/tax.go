package taxes

import (
	"math"

	"github.com/pkg/errors"
)

type TaxBracket struct {
	Threshold  float64
	Percentage float64
}

// CalculateTax takes your income (float64) and the configuration ([]TaxBracket)
// to calculate the tax amount. It returns the computed tax and any error encountered.
func CalculateTax(income float64, configuration []TaxBracket) (float64, error) {

	if income < 0 {
		return 0.0, errors.New("income cannot be lower than 0.0!")
	}

	if len(configuration) < 2 {
		return 0.0, errors.New("configuration slice must have at least two elements")
	}

	if configuration[0].Threshold != 0.0 {
		return 0.0, errors.New("the first element of the configuration slice should have threshold set to 0")
	}

	remainingAmount := income

	tax, remainingAmount := calculateTaxAmount(remainingAmount, configuration)

	if remainingAmount != 0 {
		return 0.0, errors.New("calculation error")
	}

	return tax, nil
}

func calculateTaxAmount(remainingAmount float64, taxBrackets []TaxBracket) (float64, float64) {
	var currentAmount float64
	var tax float64

	for i := 1; i < len(taxBrackets); i++ {
		percentage := taxBrackets[i-1].Percentage

		currentAmount = math.Min(taxBrackets[i].Threshold-taxBrackets[i-1].Threshold, remainingAmount)
		if percentage > 0.0 {
			tax += percentage * currentAmount
		}

		remainingAmount -= currentAmount
	}

	if remainingAmount != 0 {
		percentage := taxBrackets[len(taxBrackets)-1].Percentage // get the last element
		tax += percentage * remainingAmount
		remainingAmount -= remainingAmount
	}

	return tax, remainingAmount
}
