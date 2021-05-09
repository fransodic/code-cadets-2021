package TaxLibrary

import (
	"math"

	"github.com/pkg/errors"
)

type TaxBracketsConfiguration struct {
	Brackets   []float64
	Percentage []float64
}

// ConfigureTaxBrackets takes 2 arguments where the first one (brackets) represents
// income intervals and the second (taxRates) tax percentage for chosen intervals. Value at
// taxRates[i] is used in calculating as the taxRate for the following interval -> [brackets[i+1], brackets[i+2]].
// It returns a TaxBracketsConfiguration struct which represents the configuration for the
// calculation of progressive tax.
func ConfigureTaxBrackets(brackets []float64, taxRates []int) (TaxBracketsConfiguration, error) {
	var configuration TaxBracketsConfiguration

	if len(brackets) != len(taxRates) {
		return TaxBracketsConfiguration{},
			errors.New("arguments for func ConfigureTaxBrackets must have the same length")
	}

	configuration.Brackets = brackets
	configuration.Percentage = convertRatesToFloat64(taxRates)

	return configuration, nil
}

func convertRatesToFloat64(rates []int) []float64 {
	var taxRates []float64
	for _, percentage := range rates {
		taxRates = append(taxRates, float64(percentage)/100)
	}
	return taxRates
}

// CalculateTax takes your income (float64) and the configuration returned by ConfigureTaxBrackets
// and calculates the tax amount. It returns the computed tax and any error encountered.
func CalculateTax(income float64, configuration TaxBracketsConfiguration) (float64, error) {

	if income < 0 {
		return 0.0, errors.New("income cannot be lower than 0.0 HRK")
	}

	leftAmount := income

	brackets := configuration.Brackets
	taxRates := configuration.Percentage

	tax, leftAmount := calculateTaxAmount(leftAmount, brackets, taxRates)

	if leftAmount != 0 {
		return 0.0, errors.New("calculation error")
	}

	return tax, nil
}

func calculateTaxAmount(leftAmount float64, brackets []float64, taxRate []float64) (float64, float64) {
	var currentAmount float64
	var tax float64

	for i := 1; i < len(brackets); i++ {
		currentAmount = math.Min(brackets[i]-brackets[i-1], leftAmount)
		if taxRate[i-1] > 0.0 {
			tax += taxRate[i-1] * currentAmount
		}
		leftAmount -= currentAmount
	}

	if leftAmount != 0 {
		tax += taxRate[len(taxRate)-1] * leftAmount
		leftAmount -= leftAmount
	}

	return tax, leftAmount
}
