package main

import "testing"

type TestCase struct {
	address string
	expected string
}

func TestSolutionDefangIPaddr(t *testing.T) {
	cases := []TestCase{
		{
			address:  "1.1.1.1",
			expected: "1[.]1[.]1[.]1",
		},
		{
			address:  "255.100.50.0",
			expected: "255[.]100[.]50[.]0",
		},
	}


	for idx, item := range cases {
		ans := defangIPaddr(item.address)

		if ans != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, ans, item.expected)
		}
	}
}
