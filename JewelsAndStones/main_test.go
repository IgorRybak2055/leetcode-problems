package main

import (
	"testing"
)

type TestCase struct {
	J        string
	S        string
	expected int
}

func TestSolutionFirstUnique(t *testing.T) {
	testCases := []TestCase{
		{
			J:        "aA",
			S:        "aAAbbbb",
			expected: 3,
		},
		{
			J:        "z",
			S:        "ZZ",
			expected: 0,
		},

	}

	for idx, item := range testCases {
		ans := numJewelsInStones(item.J, item.S)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, ans, item.expected)
		}
	}
}
