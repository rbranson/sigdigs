package sigdigs

import "testing"

func TestCountSignificantDigits(t *testing.T) {
	suite := []struct {
		input    string
		expected int
	}{
		{
			input:    "1234",
			expected: 4,
		},
		{
			input:    "1234000",
			expected: 4,
		},
		{
			input:    "0001234",
			expected: 4,
		},
		{
			input:    "1.234",
			expected: 4,
		},
		{
			input:    "0001.234",
			expected: 4,
		},
		{
			input:    "1.234000",
			expected: 7,
		},
		{
			input:    "1.234e3",
			expected: 4,
		},
		{
			input:    "1.234E3",
			expected: 4,
		},
		{
			input:    "-1234",
			expected: 4,
		},
		{
			input:    "+1234",
			expected: 4,
		},
		{
			input:    "1234.0",
			expected: 5,
		},
	}

	for _, testCase := range suite {
		t.Run(testCase.input, func(t *testing.T) {
			got := CountSignificantDigits(testCase.input)
			if got != testCase.expected {
				t.Errorf("CountSignificantDigits(%s) = %d; want %d", testCase.input, got, testCase.expected)
			}
		})
	}
}
