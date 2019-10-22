package main

import (
	"math"
	"testing"
)

type Case struct {
	str      string
	expected int
}

func TestSolutionStringToInteger(t *testing.T) {
	cases := []Case {
		{
			str:      "42",
			expected: 42,
		},
		{
			str:      "     -42",
			expected: -42,
		},
		{
			str:      "words and 987",
			expected: 0,
		},
		{
			str:      "987 and words",
			expected: 987,
		},
		{
			str:      "-91283472332",
			expected: math.MinInt32,
		},
		{
			str:      "only text here",
			expected: 0,
		},
	}

	for idx, item := range cases {
		answer := myAtoi(item.str)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
