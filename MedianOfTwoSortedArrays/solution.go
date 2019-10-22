package main

import (
	"fmt"
	"sort"
)

func main()  {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	middle := len(nums1)/2
	if len(nums1)%2 == 0 {
		return float64(nums1[middle] + nums1[middle - 1]) / 2
	}
	return float64(nums1[middle])
}
