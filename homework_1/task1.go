package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func parseProgramArguments(start, end *int) {
	flag.IntVar(start, "start", 1, "Value (inclusive) from which to start the FizzBuzz game")
	flag.IntVar(end, "end", 10, "Value (inclusive) to end the FizzBuzz game")

	flag.Parse()
}

func playFizzBuzz(start int, end int) []string {
	var toPrint []string

	for i := start; i <= end; i++ {

		if i%3 == 0 {
			toPrint = append(toPrint, "Fizz")
		} else if i%5 == 0 {
			toPrint = append(toPrint, "Buzz")
		} else if i%15 == 0 {
			toPrint = append(toPrint, "Fizz Buzz")
		} else {
			toPrint = append(toPrint, strconv.Itoa(i))
		}
	}
	return toPrint
}

func print(slice []string) {
	fmt.Println(strings.Join(slice, " "))
}

func main() {
	var start, end int

	parseProgramArguments(&start, &end)

	sliceToPrint := playFizzBuzz(start, end)

	print(sliceToPrint)
}
