package dynamicprogramming

import "math"

// https://leetcode.com/explore/interview/card/google/64/dynamic-programming-4/3085/
// func maxSubArray(nums []int) int {
// 	max_sum, curr_sum := nums[0], 0

// 	for _, num := range nums {
// 		curr_sum += num

// 		max_sum = max(max_sum, curr_sum)

// 		if curr_sum < 0 {
// 			curr_sum = 0
// 		}
// 	}

// 	return max_sum
// }

func maxSubArray(nums []int) int {
	return findBestArray(nums, 0, len(nums)-1)
}

func findBestArray(nums []int, left, right int) int {
	if left > right {
		return math.MinInt
	}

	mid := left + (right-left)/2
	curr, best_left_sum, best_right_sum := 0, 0, 0

	for i := mid - 1; i >= left; i-- {
		curr += nums[i]
		best_left_sum = max(best_left_sum, curr)
	}

	curr = 0
	for i := mid + 1; i <= right; i++ {
		curr += nums[i]
		best_right_sum = max(best_right_sum, curr)
	}

	best_combined_sum := best_left_sum + nums[mid] + best_right_sum
	left_sum := findBestArray(nums, left, mid-1)
	right_sum := findBestArray(nums, mid+1, right)

	return max(best_combined_sum, max(left_sum, right_sum))
}
