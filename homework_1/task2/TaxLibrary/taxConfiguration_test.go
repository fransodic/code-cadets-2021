package TaxLibrary_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/task2/TaxLibrary"
)

func TestTaxConfiguration(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualConfiguration, actualErr := TaxLibrary.ConfigureTaxBrackets(tc.brackets, tc.percentages)

			if tc.expectingConfError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualConfiguration, ShouldResemble, tc.expectedConfiguration)
			}

		})
	}
}
