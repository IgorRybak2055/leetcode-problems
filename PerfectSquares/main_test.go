package main

import "testing"

type Case struct {
	num      int
	expected int
}

func TestSolutionIsPalindromeNumber(t *testing.T) {
	cases := []Case{
		{
			num:      12,
			expected: 3,
		},
		{
			num:      13,
			expected: 2,
		},
		{
			num:      29,
			expected: 2,
		},
		{
			num:      19,
			expected: 3,
		},
	}

	for idx, item := range cases {
		answer := numSquares(item.num)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
