package main

import "testing"

type Case struct {
	str        string
	expected string
}

func TestSolutionLongestPalindrome(t *testing.T) {
	cases := []Case{
		{
			str:      "babad",
			expected: "bab",
		},
		{
			str:      "bbbbb",
			expected: "bbbbb",
		},
		{
			str:      "a",
			expected: "a",
		},
		{
			str:      "craracr",
			expected: "rar",
		},
		{
			str:      "aabccaa",
			expected: "aa",
		},
	}

	for idx, item := range cases {
		answer := longestPalindrome(item.str)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
