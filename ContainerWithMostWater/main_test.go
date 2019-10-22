package main

import "testing"

type Case struct {
	height []int
	expected int
}

func TestSolutionMaxAreaWater(t *testing.T) {
	cases := []Case {
		{
			height:   []int{1,8,6,2,5,4,8,3,7},
			expected: 49,
		},
		{
			height:   []int{1,3,6,5,2,4,3,8,2,7},
			expected: 42,
		},
		{
			height:   []int{1,1,3,3,1,1},
			expected: 5,
		},
		{
			height:   []int{1,1,7,6,1,1},
			expected: 6,
		},
		{
			height:   []int{1,2,3,4,5,6,7,8,9},
			expected: 20,
		},
		{
			height:   []int{9,1,2,3,4,5,6,7,8,9},
			expected: 81,
		},

	}

	for idx, item := range cases {
		var answer =  maxArea(item.height)
		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
