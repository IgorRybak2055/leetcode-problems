package main

import (
	"reflect"
	"testing"
)

type Case struct {
	nums []int
	target int
	expected []int
}

func TestSolutionTwoSum(t *testing.T) {
	cases := []Case{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0,1},
		},
		{
			nums:     []int{15, 8, 1, 12},
			target:   16,
			expected: []int{0,2},
		},
		{
			nums:     []int{2, 7, 10, 4},
			target:   14,
			expected: []int{2,3},
		},
		{
			nums:     []int{2, 2, 3, 2},
			target:   4,
			expected: []int{0,1},
		},
	}

	for idx, item := range cases {
		answer := twoSum(item.nums, item.target)

		if !reflect.DeepEqual(answer, item.expected) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
