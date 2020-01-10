package main

import (
	"testing"
)

type TestCase struct {
	strs []string
	prefix string
}

func TestSolutionLongestCommonPrefix(t *testing.T) {
	cases := []TestCase{
		{
			strs:   []string{"flower","flow","flight"},
			prefix: "fl",
		},
		{
			strs:   []string{"supera","superAndA","superA1"},
			prefix: "super",
		},
		{
			strs:   []string{"c","c"},
			prefix: "c",
		},
		{
			strs:   []string{"abc","abcdfr","abcre","abctest","abc1043"},
			prefix: "abc",
		},
		{
			strs:   []string{""},
			prefix: "",
		},
		{
			strs:   []string{"Flow","flow"},
			prefix: "",
		},
	}

	for idx, item := range cases {
		answer := longestCommonPrefix(item.strs)

		if item.prefix != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.prefix)
		}
	}
}
