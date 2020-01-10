package main

import (
	"testing"
)

type TestCase struct {
	text     string
	expected bool
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase{
		{
			text:     "a man, a plan, a canal: panama",
			expected: true,
		},
		{
			text:     "race a car",
			expected: false,
		},
		{
			text:     "",
			expected: true,
		},
		{
			text:     "шалаш",
			expected: true,
		},
		{
			text:     "not palindrome",
			expected: false,
		},
	}

	for idx, item := range cases {
		answer := isPalindrome(item.text)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer, item.expected)
		}
	}
}
