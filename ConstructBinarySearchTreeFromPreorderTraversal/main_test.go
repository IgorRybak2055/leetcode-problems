package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	preopder []int
	expected *TreeNode
}

func TestSolutionConstructBinaryTree(t *testing.T) {
	cases := []TestCase {
		{
			preopder: []int{8,5,1,7,10,12},
			expected: &TreeNode{
				Val:   8,
				Left:  &TreeNode{
					Val:   5,
					Left:  &TreeNode{
						Val:   1,
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
					Val:   10,
					Left:  nil,
					Right: &TreeNode{
						Val:   12,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}


	for idx, item := range cases {
		answer := bstFromPreorder(item.preopder)

		if !reflect.DeepEqual(item.expected, answer) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v",idx, answer, item.expected)
		}
	}

}
