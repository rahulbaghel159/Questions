package dynamicprogramming

import "math"

// https://leetcode.com/explore/interview/card/google/64/dynamic-programming-4/3089/
func splitArray(nums []int, k int) int {
	prefix_sum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix_sum[i+1] = prefix_sum[i] + nums[i]
	}

	memo := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = math.MaxInt
		}
	}

	return getMinimumLargestSplitSum(memo, prefix_sum, 0, k)
}

func getMinimumLargestSplitSum(memo [][]int, prefix_sum []int, curr_index, subarray_count int) int {
	if memo[curr_index][subarray_count] != math.MaxInt {
		return memo[curr_index][subarray_count]
	}

	n := len(prefix_sum) - 1

	if subarray_count == 1 {
		memo[curr_index][subarray_count] = prefix_sum[n] - prefix_sum[curr_index]
		return memo[curr_index][subarray_count]
	}

	min_sum := math.MaxInt
	for i := curr_index; i <= n-subarray_count; i++ {
		first_split_sum := prefix_sum[i+1] - prefix_sum[curr_index]

		largest_split_sum := max(first_split_sum, getMinimumLargestSplitSum(memo, prefix_sum, i+1, subarray_count-1))

		min_sum = min(min_sum, largest_split_sum)

		if first_split_sum > min_sum {
			break
		}
	}

	memo[curr_index][subarray_count] = min_sum

	return memo[curr_index][subarray_count]
}
