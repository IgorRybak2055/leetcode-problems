package main

import (
	"testing"
)

type TestCase struct {
	gas      []int
	cost     []int
	expected int
}

func TestSolutionFirstUnique(t *testing.T) {
	testCases := []TestCase{
		{
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			gas:      []int{2,3,4},
			cost:     []int{3,4,3},
			expected: -1,
		},
	}

	for idx, item := range testCases {
		ans := canCompleteCircuit(item.gas, item.cost)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, ans, item.expected)
		}
	}
}
