package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"code-cadets-2021/homework_1/fizzBuzz"
)

func parseProgramArguments(start, end *int) {
	flag.IntVar(start, "start", 1, "Value (inclusive) from which to start the FizzBuzz game")
	flag.IntVar(end, "end", 10, "Value (inclusive) to end the FizzBuzz game")

	flag.Parse()
}

func printResult(slice []string) {
	fmt.Println(strings.Join(slice, " "))
}

func main() {
	var start, end int

	parseProgramArguments(&start, &end)

	sliceToPrint, err := fizzBuzz.PlayFizzBuzz(start, end)
	if err != nil {
		log.Fatal(err)
	}

	printResult(sliceToPrint)
}
