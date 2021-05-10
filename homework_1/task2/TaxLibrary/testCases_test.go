package TaxLibrary_test

import "code-cadets-2021/homework_1/task2/TaxLibrary"

type testCase struct {
	income      float64
	brackets    []float64
	percentages []int

	expectedConfiguration TaxLibrary.TaxBracketsConfiguration
	expectedOutput        float64
	expectingConfError    bool
	expectingCalcError    bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			income:      7000,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20, 30},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 5000, 10000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30},
			},
			expectedOutput:     800,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      15000,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20, 30},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 5000, 10000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30},
			},
			expectedOutput:     2900,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      102000,
			brackets:    []float64{0, 1000, 10000, 50000, 100000},
			percentages: []int{0, 10, 20, 30, 40},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 10000, 50000, 100000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30, 0.40},
			},
			expectedOutput:     24700,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      58000,
			brackets:    []float64{0, 1000, 10000, 50000, 100000},
			percentages: []int{0, 10, 20, 30, 40},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 10000, 50000, 100000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30, 0.40},
			},
			expectedOutput:     11300,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      0,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20, 30},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 5000, 10000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30},
			},
			expectedOutput:     0,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      1000,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20, 30},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 5000, 10000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30},
			},
			expectedOutput:     0,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      999,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20, 30},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 5000, 10000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30},
			},
			expectedOutput:     0,
			expectingConfError: false,
			expectingCalcError: false,
		},
		{
			income:      -10000,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20, 30},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{
				Brackets:   []float64{0, 1000, 5000, 10000},
				Percentage: []float64{0.0, 0.10, 0.20, 0.30},
			},
			expectedOutput:     0.0,
			expectingConfError: false,
			expectingCalcError: true,
		},
		{
			income:      100000,
			brackets:    []float64{0, 1000, 5000, 10000},
			percentages: []int{0, 10, 20},

			expectedConfiguration: TaxLibrary.TaxBracketsConfiguration{},
			expectedOutput:        0.0,
			expectingConfError:    true,
			expectingCalcError:    true,
		},
	}
}
