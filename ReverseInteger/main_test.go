package main

import (
	"testing"
)

type TestCase struct {
	num      int
	expected int
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase{
		{
			num:      123,
			expected: 321,
		},
		{
			num:      -123,
			expected: -321,
		},
		{
			num:      120,
			expected: 21,
		},
		{
			num:      5,
			expected: 5,
		},
	}

	for idx, item := range cases {
		answer := reverse(item.num)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer, item.expected)
		}
	}

}
