package TaxLibrary_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/task2/TaxLibrary"
)

func TestTaxCalculation(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			conf, err := checkConfiguration(tc)
			if err == nil {
				checkTaxCalculation(tc, conf)
			}

		})
	}
}

func checkConfiguration(tc testCase) (TaxLibrary.TaxBracketsConfiguration, error) {
	actualConfiguration, actualErr := TaxLibrary.ConfigureTaxBrackets(tc.brackets, tc.percentages)

	if tc.expectingConfError {
		So(actualErr, ShouldNotBeNil)
	} else {
		So(actualErr, ShouldBeNil)
		So(actualConfiguration, ShouldResemble, tc.expectedConfiguration)
	}

	return actualConfiguration, actualErr
}

func checkTaxCalculation(tc testCase, actualConfiguration TaxLibrary.TaxBracketsConfiguration) {
	actualOutput, actualErr := TaxLibrary.CalculateTax(tc.income, actualConfiguration)

	if tc.expectingCalcError {
		So(actualErr, ShouldNotBeNil)
	} else {
		So(actualErr, ShouldBeNil)
		So(actualOutput, ShouldResemble, tc.expectedOutput)
	}
}
