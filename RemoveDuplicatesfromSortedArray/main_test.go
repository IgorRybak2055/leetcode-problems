package main

import (
	"testing"
)

type TestCase struct {
	nums []int
	expected int
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase {
		{
			nums:     []int{1,1,2},
			expected: 2,
		},
		{
			nums:     []int{0,0,1,1,1,2,2,3,3,4},
			expected: 5,
		},
		{
			nums:     []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
			expected: 1,
		},
		{
			nums:     []int{},
			expected: 0,
		},
	}


	for idx, item := range cases {
		answer := removeDuplicates(item.nums)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
