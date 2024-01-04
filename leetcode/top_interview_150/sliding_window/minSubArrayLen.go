package slidingwindow

// https://leetcode.com/problems/minimum-size-subarray-sum
func minSubArrayLen(target int, nums []int) int {
	left, right, minLen, sum, found := 0, 0, len(nums), 0, false

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			found = true
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if !found {
		minLen = 0
	}

	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
