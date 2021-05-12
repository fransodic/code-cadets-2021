package fizzbuzz_test

type testCase struct {
	inputStart int
	inputEnd   int

	expectedOutput []string
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputStart: 1,
			inputEnd:   10,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 10,
			inputEnd:   20,

			expectedOutput: []string{"Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 1,
			inputEnd:   25,

			expectedOutput: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz"},
			expectingError: false,
		},
		{
			inputStart: 12,
			inputEnd:   5,

			expectedOutput: nil,
			expectingError: true,
		},
		{
			inputStart: 5,
			inputEnd:   2,

			expectingError: true,
		},
		{
			inputStart: 10,
			inputEnd:   5,

			expectingError: true,
		},
		{
			inputStart: -5,
			inputEnd:   10,

			expectingError: true,
		},
		{
			inputStart: 5,
			inputEnd:   -10,

			expectingError: true,
		},
		{
			inputStart: 0,
			inputEnd:   -10,

			expectingError: true,
		},
		{
			inputStart: 1,
			inputEnd:   0,

			expectingError: true,
		},
	}
}
