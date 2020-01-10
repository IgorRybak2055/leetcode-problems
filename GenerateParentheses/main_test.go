package main

import (
	"reflect"
	"testing"
)

type Case struct {
	num      int
	expected []string
}

func TestSolutionIntToRoman(t *testing.T) {
	cases := []Case{
		{
			num: 3,
			expected: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			num:      1,
			expected: []string{"()"},
		},
		{
			num:      0,
			expected: []string{""},
		},
		{
			num:      2,
			expected: []string{"(())", "()()"},
		},
	}

	for idx, item := range cases {
		answer := generateParenthesis(item.num)

		if  !reflect.DeepEqual(item.expected, answer) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer, item.expected)
		}
	}
}
