package main

import "fmt"

func main()  {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

func maxArea(height []int) int  {
	var maxArea, start, end = 0, 0, len(height) - 1

	for start != end {
		currArea := min(height[start], height[end]) * (end-start)
		if currArea > maxArea {
			maxArea = currArea
		}

		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}