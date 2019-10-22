package main

import "testing"

type Case struct {
	s        string
	p        string
	expected bool
}

func TestSolutionIsMatch(t *testing.T) {
	cases := []Case{
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "aa",
			p:        "a.",
			expected: true,
		},
		{
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
		},
		{
			s:        "mississippi",
			p:        "mis*is*.p*.",
			expected: true,
		},
	}

	for idx, item := range cases {
		answer := isMatch(item.s, item.p)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
