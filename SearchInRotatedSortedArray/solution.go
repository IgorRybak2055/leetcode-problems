package main

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	search(nums, 0)
}

func search(nums []int, target int) int {
	var (
		l, m int
		r = len(nums) - 1
	)

	for l <= r {
		m = l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < nums[l] && (target >= nums[l] || target < nums[m]) {
			r = m - 1
		} else if nums[m] >= nums[l] && (target < nums[l] || target > nums[m]) {
			l = m + 1
		} else if nums[m] < nums[l] {
			l = m + 1
		} else if nums[m] >= nums[l] {
			r = m - 1
		}
	}
	return -1
}
