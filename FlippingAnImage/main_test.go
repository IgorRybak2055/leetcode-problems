package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	A        [][]int
	expected [][]int
}

func TestSolutionFirstUnique(t *testing.T) {
	testCases := []TestCase {
		{
			A:        [][]int{{1,1,0},{1,0,1},{0,0,0}},
			expected: [][]int{{1,0,0},{0,1,0},{1,1,1}},
		},
		{
			A:        [][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}},
			expected: [][]int{{1,1,0,0},{0,1,1,0},{0,0,0,1},{1,0,1,0}},
		},
	}

	for idx, item := range testCases {
		ans := 	flipAndInvertImage(item.A)

		if !reflect.DeepEqual(ans, item.expected) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, ans, item.expected)
		}
	}
}
