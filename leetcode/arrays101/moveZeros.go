package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/searching-for-items-in-an-array/3157/
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
