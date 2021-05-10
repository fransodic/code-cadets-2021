package fizzBuzz

import (
	"strconv"

	"github.com/pkg/errors"
)

func PlayFizzBuzz(start int, end int) ([]string, error) {

	if start <= 0 {
		return nil, errors.New("start should be greater than 0")
	}

	if end <= 0 {
		return nil, errors.New("end should be greater than 0")
	}

	if start > end {
		return nil, errors.New("start is greater than end")
	}

	var toPrint []string

	for i := start; i <= end; i++ {

		if i%15 == 0 {
			toPrint = append(toPrint, "FizzBuzz")
		} else if i%5 == 0 {
			toPrint = append(toPrint, "Buzz")
		} else if i%3 == 0 {
			toPrint = append(toPrint, "Fizz")
		} else {
			toPrint = append(toPrint, strconv.Itoa(i))
		}
	}
	return toPrint, nil
}
