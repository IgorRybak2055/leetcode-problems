package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	// for travis CI
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	i := 0
	for i < len(preorder) && preorder[i] <= preorder[0] {
		i++
	}

	return &TreeNode{
		Val: preorder[0],
		Left: bstFromPreorder(preorder[1:i]),
		Right: bstFromPreorder(preorder[i:]),

	}
}
