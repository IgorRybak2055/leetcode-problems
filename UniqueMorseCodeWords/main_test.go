package main

import (
	"testing"
)

type TestCase struct {
	words    []string
	expected int
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase{
		{
			words:    []string{"gin", "zen", "gig", "msg"},
			expected: 2,
		},
		{
			words:    []string{"sos", "car", "cat"},
			expected: 3,
		},
		{
			words:    []string{"", ""},
			expected: 1,
		},
	}

	for idx, item := range cases {
		answer := uniqueMorseRepresentations(item.words)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer, item.expected)
		}
	}
}
