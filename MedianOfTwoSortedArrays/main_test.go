package main

import (
	"testing"
)

type Case struct {
	nums1 []int
	nums2 []int
	expected float64
}

func TestSolutionFindMedianSortedArrays(t *testing.T) {
	cases := []Case{
		{
			nums1:     []int{1, 3},
			nums2:   []int{2},
			expected: 2.0,
		},
		{
			nums1:     []int{1, 2},
			nums2:   []int{3,4},
			expected: 2.5,
		},
		{
			nums1:     []int{-5, 3, 6, 12, 15},
			nums2:   []int{-12, -10, -6, -3, 4, 10},
			expected: 3,
		},
		{
			nums1:     []int{2, 3, 5, 8},
			nums2:   []int{10, 12, 14, 16, 18, 20},
			expected: 11,
		},
		{
			nums1:     []int{1},
			nums2:   []int{},
			expected: 1,
		},
	}

	for idx, item := range cases {
		answer := findMedianSortedArrays(item.nums1, item.nums2)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
