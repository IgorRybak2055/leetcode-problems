package main

import "testing"

type TestCase struct {
	nums     []int
	expected int
}

func TestSolutionFindDuplicateNumber(t *testing.T) {
	testCases := []TestCase{
		{
			nums:     []int{1,3,4,2,2},
			expected: 2,
		},
		{
			nums:     []int{3,1,3,4,2},
			expected: 3,
		},
		{
			nums:     []int{5,1,3,4,2},
			expected: -1,
		},
	}

	for idx, item := range testCases {
		ans := findDuplicate(item.nums)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, ans, item.expected)
		}
	}

}
