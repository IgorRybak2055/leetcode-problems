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
			nums: []int{3,2,1,6,0,5},
			expected: &TreeNode{
				Val:   6,
				Left:  &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}


	for idx, item := range cases {
		answer := constructMaximumBinaryTree(item.nums)

		if !reflect.DeepEqual(item.expected, answer) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
