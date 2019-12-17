package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	digits string
	expected []string
}

func TestSolutionLetterCombinations(t *testing.T) {
	testCases := []TestCase{
		{
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits:   "234",
			expected: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh",
				"bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
		{
			digits:   "1",
			expected: []string{},
		},
		{
			digits:   "19",
			expected: []string(nil),
		},
		{
			digits:   "9",
			expected: []string{"w", "x", "y", "z"},
		},
	}

	for idx, item := range testCases {
		ans := letterCombinations(item.digits)

		if !reflect.DeepEqual(ans, item.expected) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, ans, item.expected)
		}
	}
}
