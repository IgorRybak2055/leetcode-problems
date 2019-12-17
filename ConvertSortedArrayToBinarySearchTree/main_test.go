package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums []int
	expected *TreeNode
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase {
		{
			nums: []int{-10,-3,0,5,9},
			expected: &TreeNode{
				Val:   0,
				Left:  &TreeNode{
					Val:   -3,
					Left:  &TreeNode{
						Val:   -10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   9,
					Left:  &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}


	for idx, item := range cases {
		answer := sortedArrayToBST(item.nums)

		if !reflect.DeepEqual(item.expected, answer) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
