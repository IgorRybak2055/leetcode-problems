package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	// for travis CI
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2

	return &TreeNode{
		nums[mid],
		sortedArrayToBST(nums[:mid]),
		sortedArrayToBST(nums[mid+1:]),
	}
}
