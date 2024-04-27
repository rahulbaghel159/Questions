package dynamicprogramming

// https://leetcode.com/explore/interview/card/google/64/dynamic-programming-4/451/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
		if i+1 < len(s) && s[i] == s[i+1] {
			dp[i][i+1] = 2
			dp[i+1][i] = 2
		}
	}

	for diff := 2; diff < len(s); diff++ {
		for i := 0; i < len(s)-diff; i++ {
			j := i + diff
			if s[i] == s[j] && dp[i+1][j-1] != 0 {
				dp[i][j] = j - i + 1
			}
		}
	}

	start, end, max_len := 0, 0, 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if max_len < dp[i][j] {
				start, end = i, j
				max_len = dp[i][j]
			}
		}
	}

	if end+1 < len(s) {
		return s[start : end+1]
	}

	return s[start:]
}
