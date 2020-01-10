package main

import (
	"log"
)

func main() {
	nums := []int{5, 7, 8, 10, 10}
	log.Println(searchRange(nums, 8))
}

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}
	start, end := 0, numsLen-1
	var mid int

	// search left
	for start < end {
		if mid = start + (end-start)/2; nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if start < 0 || end >= numsLen || nums[start] != target {
		return []int{-1, -1}
	}
	left := end

	// search right
	start, end = left, numsLen-1
	for start < end {
		if mid = end - (end-start)/2; nums[mid] <= target {
			start = mid
		} else {
			end = mid - 1
		}
	}
	right := start

	return []int{left, right}
}
