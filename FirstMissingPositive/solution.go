package main

import "log"

func main() {
	// var nums = []int{4, 1, 2, 3}
	var nums = []int{0, 2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	log.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	l := len(nums)
	
	for i := 0; i < l; i++ {
		if nums[i] >= 1 && nums[i] <= l && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l + 1
}
