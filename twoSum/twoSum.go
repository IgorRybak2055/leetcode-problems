package twoSum

func twoSum(nums []int, target int) []int {
	sum := make([]int, 0)
	for ind1 := 0; ind1 < len(nums); ind1++ {
		find := target - nums[ind1]
		for ind2 := ind1 + 1; ind2 < len(nums); ind2++  {
			if nums[ind2] == find {
				sum= append(sum, ind1, ind2)
				return sum
			}
		}
	}
	return sum
}

