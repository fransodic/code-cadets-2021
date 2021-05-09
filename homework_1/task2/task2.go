package main

import (
	"code-cadets-2021/homework_1/task2/TaxLibrary"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {

	income := getFromStdin()

	configuration, err := TaxLibrary.ConfigureTaxBrackets([]float64{0, 1000, 5000, 10000}, []int{0, 10, 20, 30})

	tax, err := TaxLibrary.CalculateTax(income, configuration)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Ukupni porez za %v HRK prihoda iznosi %v HRK.", income, tax)
	}
}

func getFromStdin() float64 {
	var income float64

	fmt.Print("Unesite prihod: ")

	_, err := fmt.Scanf("%f", &income)
	if err != nil {
		errors.New("trouble reading from stdin")
	}

	return income
}
