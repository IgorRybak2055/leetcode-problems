package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	// for travis CI
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if root.Val < L {
		return rangeSumBST(root.Right,L,R)
	}

	if root.Val > R {
		return rangeSumBST(root.Left,L,R)
	}

	return root.Val + rangeSumBST(root.Right,L,R) + rangeSumBST(root.Left,L,R)
}
