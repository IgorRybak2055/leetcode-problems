package main

import "testing"

type Case struct {
	str      string
	numsRows int
	expected string
}

func TestSolutionZigZagConvert(t *testing.T) {
	cases := []Case{
		{
			str:      "PAYPALISHIRING",
			numsRows: 3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			str:      "PAYPALISHIRING",
			numsRows: 4,
			expected: "PINALSIGYAHRPI",
		},
		{
			str:      "AB",
			numsRows: 1,
			expected: "AB",
		},
		{
			str:      "ABCD",
			numsRows: 3,
			expected: "ABDC",
		},
		{
			str:      "STREET",
			numsRows: 9,
			expected: "STREET",
		},
		{
			str:      "STRETE",
			numsRows: 5,
			expected: "STREET",
		},
	}

	for idx, item := range cases {
		answer := convert(item.str, item.numsRows)

		if answer != item.expected {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}
}
