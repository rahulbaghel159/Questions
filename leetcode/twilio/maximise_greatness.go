package twilio

import "sort"

// https://leetcode.com/problems/maximize-greatness-of-an-array/
func maximizeGreatness(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res++
		}
	}

	return res
}
