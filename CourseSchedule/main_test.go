package main

import "testing"

type TestCase struct {
	numCourses int
	prerequisites [][]int
	expected bool
}

func TestSolutionCourseSchedule(t *testing.T) {
	cases := []TestCase{
		{
			numCourses:    2,
			prerequisites: [][]int{{1,0}},
			expected:      true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1,0}, {0,1}},
			expected:      false,
		},
	}

	for idx, item := range cases {
		answer := CanFinish(item.numCourses, item.prerequisites)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
