package main

import "testing"

type Case struct {
	num      int
	expected bool
}

func TestSolutionIsPalindromeNumber(t *testing.T) {
	cases := []Case{
		{
			num:      121,
			expected: true,
		},
		{
			num:      -121,
			expected: false,
		},
		{
			num:      0,
			expected: true,
		},
		{
			num:      12,
			expected: false,
		},
		{
			num:      10001,
			expected: true,
		},
		{
			num:      1000,
			expected: false,
		},

	}

	for idx, item := range cases {
		answer := isPalindrome(item.num)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
