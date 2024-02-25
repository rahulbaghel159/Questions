package multidp

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	dp, ans := make([][]bool, n), make([]int, 2)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans[0], ans[1] = i, i+1
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff
			if s[j] == s[i] && dp[i+1][j-1] {
				dp[i][j] = true
				ans[0], ans[1] = i, j
			}
		}
	}

	return s[ans[0] : ans[1]+1]
}
