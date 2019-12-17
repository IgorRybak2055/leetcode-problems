package main

import "testing"

type Case struct {
	grid 	[][]int
	expected int
}

func TestSolutionLongestPalindrome(t *testing.T) {
	cases := []Case{
		{
			grid:     [][]int{{3,0,8,4},{2,4,5,7},{9,2,6,3},{0,3,1,0}},
			expected: 35,
		},
	}

	for idx, item := range cases {
		answer := maxIncreaseKeepingSkyline(item.grid)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
