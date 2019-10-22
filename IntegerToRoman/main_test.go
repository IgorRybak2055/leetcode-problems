package main

import "testing"

type Case struct {
	num      int
	expected string
}

func TestSolutionIntToRoman(t *testing.T) {
	cases := []Case{
		{
			num:      3,
			expected: "III",
		},
		{
			num:      11,
			expected: "XI",
		},
		{
			num:      4,
			expected: "IV",
		},
		{
			num:      28,
			expected: "XXVIII",
		},
		{
			num:      98,
			expected: "XCVIII",
		},
		{
			num:      1994,
			expected: "MCMXCIV",
		},
	}

	for idx, item := range cases {
		answer := intToRoman(item.num)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
