package main

import (
	"flag"
	"fmt"
	"log"

	"code-cadets-2021/homework_1/task2/TaxLibrary"
)

func parseProgramArguments(income *float64) {
	flag.Float64Var(income, "income", 1, "Your income in $$")

	flag.Parse()
}

func main() {
	var income float64

	parseProgramArguments(&income)

	configuration := TaxLibrary.ConfigureTaxBrackets([]float64{0, 9075, 36900, 89350, 186350, 405100}, []int{0, 10, 15, 20, 25, 28, 33})

	tax, err := TaxLibrary.CalculateTax(income, configuration)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(tax)
}
