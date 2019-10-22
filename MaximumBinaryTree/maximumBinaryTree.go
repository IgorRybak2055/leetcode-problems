package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }

func main()  {
	// for travis CI
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	indOfMax := max(nums)

	node := TreeNode{Val:nums[indOfMax]}
	if indOfMax >= 1 {
		node.Left = constructMaximumBinaryTree(nums[:indOfMax])
	}

	if indOfMax + 1 < len(nums) {
		node.Right = constructMaximumBinaryTree(nums[indOfMax+1:])
	}

	return &node
}

func max(nums []int) int {
	m, retInd := 0, -1
	for ind, val := range nums{
		if m < val {
			m = val
			retInd = ind
		}
	}

	return retInd
}
