package dynamicprogramming

// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/3035/
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp, ans := make([]int, len(s)), 0
	for i := 1; i < len(dp); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-1 >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}

			ans = max(ans, dp[i])
		}
	}

	return ans
}
