package main

import "testing"

type TestCase struct {
	str string
	expected bool
}

func TestSolutionValidParentheses(t *testing.T) {
	testCases := []TestCase{
		{
			str:      "()",
			expected: true,
		},
		{
			str:      "()[]{}",
			expected: true,
		},
		{
			str:      "(]",
			expected: false,
		},
		{
			str:      "{[]}",
			expected: true,
		},
		{
			str:      "([)]",
			expected: false,
		},
		{
			str:      "(){",
			expected: false,
		},
		{
			str:      "",
			expected: true,
		},
	}

	for idx, item := range testCases{
		ans := isValid(item.str)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, ans, item.expected)
		}
	}
}


