package main

import (
	"testing"
)

type TestCase struct {
	str string
	expected string
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase {
		{
			str:      "(()())(())",
			expected: "()()()",
		},
		{
			str:      "(()())(())(()(()))",
			expected: "()()()()(())",
		},
		{
			str:      "()()",
			expected: "",
		},
	}


	for idx, item := range cases {
		answer := removeOuterParentheses(item.str)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
