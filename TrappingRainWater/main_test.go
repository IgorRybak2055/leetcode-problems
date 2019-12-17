package main

import (
	"testing"
)

type TestCase struct {
	height   []int
	expected int
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase{
		{
			height:   []int{0,1,0,2,1,0,1,3,2,1,2,1},
			expected: 6,
		},
		{
			height:   []int{1,1,1,1,1,2,1,1,1},
			expected: 0,
		},
	}

	for idx, item := range cases {
		answer := trap(item.height)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer, item.expected)
		}
	}

}
