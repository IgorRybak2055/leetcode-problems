package leetcodeProblems

func removeDuplicates(nums []int) int {
	m := make(map[int]bool)
	var a int

	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = true
			nums[a] = val
			a++
		}
	}

	nums = nums[:a]


	return len(nums)
}


