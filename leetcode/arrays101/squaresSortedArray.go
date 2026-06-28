package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3240/
func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := make([]int, len(nums))

	for i, j, k := 0, len(nums)-1, len(nums)-1; i <= j; {
		if abs(nums[i]) > abs(nums[j]) {
			result[k] = nums[i] * nums[i]
			i++
		} else {
			result[k] = nums[j] * nums[j]
			j--
		}
		k--
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
