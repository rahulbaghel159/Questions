package dynamicprogramming

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	return recursiveCheckSubarray(nums, k, 0, 0, 0, make(map[string]bool))
}

func recursiveCheckSubarray(nums []int, k, items, sum_so_far, index int, memo map[string]bool) bool {
	key := fmt.Sprint(sum_so_far) + ":" + fmt.Sprint(index)

	if _, ok := memo[key]; ok {
		return memo[key]
	}

	if sum_so_far%k == 0 && items >= 2 {
		return true
	}

	if index >= len(nums) {
		return false
	}

	memo[key] = recursiveCheckSubarray(nums, k, items+1, sum_so_far+nums[index], index+1, memo) || recursiveCheckSubarray(nums, k, 1, nums[index], index+1, memo)

	return memo[key]
}
