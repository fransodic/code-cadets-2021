package taxes

import (
	"math"

	"github.com/pkg/errors"
)

type TaxBracket struct {
	Threshold  float64
	Percentage float64
}

//
//// ConfigureTaxBrackets takes 2 arguments where the first one (brackets) represents
//// income intervals and the second (taxRates) tax percentage for chosen intervals. Value at
//// taxRates[i] is used in calculating as the taxRate for the following interval -> [brackets[i+1], brackets[i+2]].
//// It returns a TaxBracketsConfiguration struct and any error encountered.
//func ConfigureTaxBrackets(brackets []float64, taxRates []int) (TaxBracketsConfiguration, error) {
//	var configuration TaxBracketsConfiguration
//
//	if len(brackets) != len(taxRates) {
//		return TaxBracketsConfiguration{},
//			errors.New("arguments for func ConfigureTaxBrackets must have the same length")
//	}
//
//	configuration.Brackets = brackets
//	configuration.Percentage = convertRatesToFloat64(taxRates)
//
//	return configuration, nil
//}
//
//func convertRatesToFloat64(rates []int) []float64 {
//	var taxRates []float64
//	for _, percentage := range rates {
//		taxRates = append(taxRates, float64(percentage)/100)
//	}
//	return taxRates
//}

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
