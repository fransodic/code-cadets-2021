package main

import (
	"code-cadets-2021/homework_1/task2/taxes"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {

	income, err := getFromStdin()
	if err != nil {
		log.Fatal(
			err,
		)
	}

	//configuration := buildConfiguration([]float64{0, 1000, 5000, 10000}, []float64{0.0, 0.10, 0.20, 0.30})
	configuration := buildConfiguration([]float64{2000, 3000}, []float64{0.10, 0.20})

	tax, err := taxes.CalculateTax(income, configuration)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "trouble calculating tax"),
		)
	}

	fmt.Printf("Ukupni porez za vrijednost %v iznosi %v.", income, tax)

}

func buildConfiguration(thresholds []float64, percentages []float64) []taxes.TaxBracket {
	var configuration []taxes.TaxBracket

	for i := 0; i < len(thresholds); i++ {
		taxBracket := taxes.TaxBracket{
			Threshold:  thresholds[i],
			Percentage: percentages[i],
		}
		configuration = append(configuration, taxBracket)
	}

	return configuration
}

func getFromStdin() (float64, error) {
	var income float64

	fmt.Print("Unesite prihod: ")

	_, err := fmt.Scanf("%f", &income)
	if err != nil {
		return 0.0, errors.New("trouble getting income from stdin")
	}

	return income, nil
}
