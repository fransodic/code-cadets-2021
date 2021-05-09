package fizzBuzz_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/fizzBuzz"
)

func TestPlayFizzBuzzGame(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := fizzBuzz.PlayFizzBuzz(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}
