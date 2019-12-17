package main

import (
	"testing"
)

type TestCase struct {
	root *TreeNode
	L int
	R int
	expected int
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase {
		{
			root:     &TreeNode{
				Val:   10,
				Left:  &TreeNode{
					Val:   5,
					Left:  &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: &TreeNode{
						Val:   18,
						Left:  nil,
						Right: nil,
					},
				},
			},
			L:        7,
			R:        15,
			expected: 32,
		},
		{
			root:     &TreeNode{
				Val:   10,
				Left:  &TreeNode{
					Val:   5,
					Left:  &TreeNode{
						Val:   3,
						Left:  &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   15,
					Left:  &TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   18,
						Left:  nil,
						Right: nil,
					},
				},
			},
			L:        6,
			R:        10,
			expected: 23,
		},

	}


	for idx, item := range cases {
		answer := rangeSumBST(item.root, item.L, item.R)

		if item.expected != answer {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
