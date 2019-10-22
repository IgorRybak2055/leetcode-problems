package main

import (
	"reflect"
	"testing"
)

type Case struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}

func TestSolutionLongestPalindrome(t *testing.T) {
	cases := []Case{
		{
			nums1:    []int{1,2,3,0,0,0},
			m:        3,
			nums2:    []int{2,5,6},
			n:        3,
			expected: []int{1,2,2,3,5,6},
		},
	}

	for idx, item := range cases {
		merge(item.nums1, item.m, item.nums2, item.n)

		if !reflect.DeepEqual(item.nums1, item.expected) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, item.nums1, item.expected)
		}
	}

}
