package main

import (
	"log"
)

func main() {
	var (
		nums = []int{3, 2, 2, 3}
		val  = 3
	)

	log.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	ind := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ind] = nums[i]
			ind++
		}
	}

	return ind
}
