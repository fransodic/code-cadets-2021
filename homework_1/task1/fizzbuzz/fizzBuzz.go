package fizzbuzz

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

	var fizzBuzzOutput []string

	for i := start; i <= end; i++ {

		if i%15 == 0 {
			fizzBuzzOutput = append(fizzBuzzOutput, "FizzBuzz")
		} else if i%5 == 0 {
			fizzBuzzOutput = append(fizzBuzzOutput, "Buzz")
		} else if i%3 == 0 {
			fizzBuzzOutput = append(fizzBuzzOutput, "Fizz")
		} else {
			fizzBuzzOutput = append(fizzBuzzOutput, strconv.Itoa(i))
		}
	}
	return fizzBuzzOutput, nil
}
