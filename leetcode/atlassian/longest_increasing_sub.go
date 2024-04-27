package atlassian

// https://leetcode.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLength := make([]int, len(nums))
	for i := len(maxLength) - 1; i >= 0; i-- {
		maxLength[i] = 1
		if i < len(maxLength)-1 {
			for j := i + 1; j < len(maxLength); j++ {
				if nums[i] < nums[j] {
					maxLength[i] = max(maxLength[i], maxLength[j]+1)
				}
			}
		}
	}

	ans := maxLength[0]
	for i := 0; i < len(maxLength); i++ {
		ans = max(ans, maxLength[i])
	}

	return ans
}
