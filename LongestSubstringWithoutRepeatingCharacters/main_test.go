package main

import "testing"

type Case struct {
	str string
	expected int
}

func TestSolutionLengthOfLongestSubstring(t *testing.T) {

	cases := []Case{
		{
			str:      "abcabcbb",
			expected: 3,
		},
		{
			str:      "bbbbb",
			expected: 1,
		},
		{
			str:      "pwwkew",
			expected: 3,
		},
		{
			str:      "pwwkew",
			expected: 3,
		},
		{
			str:      "dvdf",
			expected: 3,
		},
		{
			str:      "aab",
			expected: 2,
		},
	}

	for idx, item := range cases {
		answer := lengthOfLongestSubstring(item.str)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
