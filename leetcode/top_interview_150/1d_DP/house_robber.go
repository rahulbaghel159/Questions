package ddp

// https://leetcode.com/problems/house-robber
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i <= n; i++ {
		if i < n {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		} else {
			dp[i] = max(dp[i-2], dp[i-1])
		}
	}

	return dp[n]
}
