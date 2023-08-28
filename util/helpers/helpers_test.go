package helpers

import "testing"

type CountCharsCase struct {
	input     string
	expectedC int
	expectedV int
}

func TestCountChars(t *testing.T) {
	cases := []CountCharsCase{
		{"hello", 3, 2},
		{"Olivia Wilson", 6, 6},
		{"Michael Jackson", 9, 5},
	}
	for _, tCase := range cases {
		resultC, resultV := CountChars(tCase.input)
		if resultC != tCase.expectedC || resultV != tCase.expectedV {
			t.Errorf("For CountChars(%s), received %d and %d, expected was %d and %d", tCase.input, resultC, resultV, tCase.expectedC, tCase.expectedV)
		}
	}
}

type ShareCommonFactorsCase struct {
	inputA   int
	inputB   int
	expected bool
}

func TestShareCommonFactors(t *testing.T) {
	cases := []ShareCommonFactorsCase{
		{3, 2, false},
		{6, 6, true},
		{9, 5, false},
		{20, 2, true},
		{50, 10, true},
		{200, 100, true},
	}
	for _, tCase := range cases {
		result := ShareCommonFactors(tCase.inputA, tCase.inputB)
		if result != tCase.expected {
			t.Errorf("For ShareCommonFactors(%d, %d), received %v, expected was %v", tCase.inputA, tCase.inputB, result, tCase.expected)
		}
	}
}
