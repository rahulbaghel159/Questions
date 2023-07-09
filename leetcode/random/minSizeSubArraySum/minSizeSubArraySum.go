package minsizesubarraysum

//https://leetcode.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	start, minLen, sum := 0, 999999999999999999, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target && start <= i {
			if sum >= target {
				minLen = min(minLen, i-start+1)
			}
			sum -= nums[start]
			start++
		}
	}

	if minLen == 999999999999999999 {
		minLen = 0
	}

	return minLen
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
