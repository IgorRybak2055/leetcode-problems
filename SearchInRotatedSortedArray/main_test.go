package main

import (
	"testing"
)

type TestCase struct {
	nums     []int
	target   int
	expected int
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase{
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: 0,
		},
	}

	for idx, item := range cases {
		answer := search(item.nums, item.target)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer, item.expected)
		}
	}

}
