package dynamicprogramming

// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/3034/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	dp := make([][]bool, len(s))
	start, end := 0, 0

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true

		if i+1 < len(s) && s[i] == s[i+1] {
			dp[i][i+1] = true
			start, end = i, i+1
		}
	}

	for diff := 2; diff < len(s); diff++ {
		for i := 0; i < len(s)-diff; i++ {
			j := i + diff
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start, end = i, j
			}
		}
	}

	return s[start : end+1]
}
