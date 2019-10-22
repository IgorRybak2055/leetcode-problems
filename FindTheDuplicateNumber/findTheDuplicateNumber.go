package main

func findDuplicate(nums []int) int {
	myMap := make(map[int]bool)

	for _, val := range nums {
		if _, ok := myMap[val]; ok {
			return val
		}
		myMap[val] = true
	}
	return -1
}
