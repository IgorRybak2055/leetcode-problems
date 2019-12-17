package main

import (
	"sort"
)

func main() {

}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var resp = make([][]int, 0, 2)

	for idx := 0; idx < len(nums) - 2; idx++ {
		secIdx, lIdx, diff := idx + 1, len(nums) - 1, -nums[idx]

		for secIdx< lIdx {
			sum := nums[secIdx] + nums[lIdx]
			switch {
			case sum < diff:
				secIdx++
			case sum > diff:
				lIdx--
			default:
				resp = append(resp, []int{nums[idx], nums[secIdx], nums[lIdx]})

				for secIdx+ 1 < lIdx && nums[secIdx] == nums[secIdx+1] {
					secIdx++
				}
				for lIdx - 1 > secIdx&& nums[lIdx] == nums[lIdx-1] {
					lIdx--
				}
				secIdx++
				lIdx--
			}
		}
		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx++
		}
	}
	return resp
}