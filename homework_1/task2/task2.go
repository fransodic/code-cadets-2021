package main

import (
	"code-cadets-2021/homework_1/task2/TaxLibrary"
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

	configuration, err := TaxLibrary.ConfigureTaxBrackets([]float64{0, 1000, 5000, 10000}, []int{0, 10, 20, 30})
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "trouble configuring the tax brackets"),
		)
	}

	tax, err := TaxLibrary.CalculateTax(income, configuration)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "trouble calculating tax"),
		)
	}

	fmt.Printf("Ukupni porez za %v HRK prihoda iznosi %v HRK.", income, tax)

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
