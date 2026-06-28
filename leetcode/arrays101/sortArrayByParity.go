package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/searching-for-items-in-an-array/3260/
func sortArrayByParity(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	k := 0

	for i := 0; i < len(nums); i++ {
		if isEven(nums[i]) {
			tmp := nums[k]
			nums[k] = nums[i]
			nums[i] = tmp
			k++
		}
	}

	return nums
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}
