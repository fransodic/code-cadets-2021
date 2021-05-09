package main

import (
	"flag"
	"log"

	"code-cadets-2021/lecture_1/07_testing/filter"
)

func main() {
	var countFrom, countTo, divisor int
	// isto kao
	//countFrom, countTo, divisor := 0, 0, 0

	flag.IntVar(&countFrom, "count-from", 1, "Value (inclusive) from which to start counting")
	flag.IntVar(&countTo, "count-to", 10, "Value (inclusive) to count to")
	flag.IntVar(&divisor, "divisor", 1, "The divisor used for filtering")

	// za isparsirati argumente
	flag.Parse()

	numbers, err := filter.GetDivisibleFromRange(countFrom, countTo, divisor)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", numbers)
}
