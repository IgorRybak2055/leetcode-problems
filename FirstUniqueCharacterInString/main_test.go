package main

import "testing"

type TestCase struct {
	s        string
	expected int
}

func TestSolutionFirstUnique(t *testing.T) {
	testCases := []TestCase {
		{
			s:        "leetcode",
			expected: 0,
		},
		{
			s:        "loveleetcode",
			expected: 2,
		},
		{
			s:        "aaaa",
			expected: -1,
		},
	}

	for idx, item := range testCases {
		ans := 	firstUniqChar(item.s)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, ans, item.expected)
		}
	}
}
