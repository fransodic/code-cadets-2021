package taxes_test

import "code-cadets-2021/homework_1/task2/taxes"

type testCase struct {
	income float64

	configuration      []taxes.TaxBracket
	expectedOutput     float64
	expectingConfError bool
	expectingError     bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			income: 7000,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{5000, 0.20},
				{10000, 0.30},
			},
			expectedOutput: 800,
			expectingError: false,
		},
		{
			income: 15000,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{5000, 0.20},
				{10000, 0.30},
			},
			expectedOutput: 2900,
			expectingError: false,
		},
		{
			income: 102000,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{10000, 0.20},
				{50000, 0.30},
				{100000, 0.40},
			},
			expectedOutput: 24700,
			expectingError: false,
		},
		{
			income: 58000,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{10000, 0.20},
				{50000, 0.30},
				{100000, 0.40},
			},
			expectedOutput: 11300,
			expectingError: false,
		},
		{
			income: 0,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{5000, 0.20},
				{10000, 0.30},
			},
			expectedOutput: 0,
			expectingError: false,
		},
		{
			income: 1000,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{5000, 0.20},
				{10000, 0.30},
			},
			expectedOutput: 0,
			expectingError: false,
		},
		{
			income: 999,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{5000, 0.20},
				{10000, 0.30},
			},
			expectedOutput: 0,
			expectingError: false,
		},
		{
			income: -10000,

			configuration: []taxes.TaxBracket{
				{0, 0.0},
				{1000, 0.10},
				{5000, 0.20},
				{10000, 0.30},
			},
			expectedOutput: 0.0,
			expectingError: true,
		},
		{
			income: 100000,

			configuration:  []taxes.TaxBracket{},
			expectedOutput: 0.0,
			expectingError: true,
		},
		{
			income: 10000,

			configuration:  []taxes.TaxBracket{{0, 0.0}},
			expectedOutput: 0.0,
			expectingError: true,
		},
		{
			income: 10000,

			configuration:  []taxes.TaxBracket{{200, 0.10}, {1000, 0.20}},
			expectedOutput: 0.0,
			expectingError: true,
		},
	}
}
