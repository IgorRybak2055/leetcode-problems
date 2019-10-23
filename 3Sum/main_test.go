package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums     []int
	expected [][]int
}

func TestSolutionThreeSum(t *testing.T) {
	testCases := []TestCase{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{-3, -2, -2, 0, 4, 5},
			expected: [][]int{{-3,-2,5}, {-2,-2,4}},
		},
		{
			nums:     []int{1,1,1},
			expected: [][]int{},
		},
		{
			nums:     []int{-2,-1,-1,1,4},
			expected: [][]int{},
		},
	}

	for idx, item := range testCases {
		answer := threeSum(item.nums)

		if !reflect.DeepEqual(answer, item.expected) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
