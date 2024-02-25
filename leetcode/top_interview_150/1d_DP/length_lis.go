package ddp

// https://leetcode.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	longest := 0
	for _, v := range dp {
		longest = max(longest, v)
	}

	return longest
}
