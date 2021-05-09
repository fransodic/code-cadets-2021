package TaxLibrary

import "math"

type TaxBracketsConfiguration struct {
	brackets   []float64
	percentage []float64
}

// ConfigureTaxBrackets takes 2 arguments where the first one (brackets) represents
// income intervals and the second (taxRates) tax percentage for chosen intervals. Value at
// taxRates[i] is used in calculating the tax for the following interval - brackets[i] - brackets[i+1].
// It returns a TaxBracketsConfiguration struct which represents the configuration for the
// calculation of progressive tax.
//
func ConfigureTaxBrackets(brackets []float64, taxRates []int) TaxBracketsConfiguration {
	var configuration TaxBracketsConfiguration

	configuration.brackets = brackets
	for _, percentage := range taxRates {
		configuration.percentage = append(configuration.percentage, float64(percentage)/100)
	}

	return configuration
}

// CalculateTax takes your income (float64) and the configuration returned by ConfigureTaxBrackets
// and calculates the tax amount. It returns the computed tax and any error encountered.
func CalculateTax(income float64, configuration TaxBracketsConfiguration) (float64, error) {
	leftAmount := income
	var tax float64

	brackets := configuration.brackets
	taxRate := configuration.percentage

	for i := 1; i < len(brackets); i++ {
		currentAmount := math.Min(brackets[i]-brackets[i-1], leftAmount)
		tax += taxRate[i] * currentAmount
		leftAmount -= currentAmount
	}

	return tax, nil
}
