package main

import (
	"testing"
)

type TestCase struct {
	haystack string
	needle   string
	expected int
}

func TestSolutionFirstUnique(t *testing.T) {
	testCases := []TestCase{
		{
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			haystack: "test",
			needle:   "dd",
			expected: -1,
		},
		{
			haystack: "Solved",
			needle:   "Solved",
			expected: 0,
		},
		{
			haystack: "default needle",
			needle:   "",
			expected: 0,
		},
	}

	for idx, item := range testCases {
		ans := strStr(item.haystack, item.needle)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, ans, item.expected)
		}
	}
}
