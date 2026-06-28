package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/searching-for-items-in-an-array/3575/
func removeElement2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
