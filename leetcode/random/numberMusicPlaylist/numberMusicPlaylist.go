package numbermusicplaylist

//https://leetcode.com/problems/number-of-music-playlists

func numMusicPlaylists(n int, goal int, k int) int {
	dp := make([][]int, goal+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	dp[0][0] = 1

	for i := 1; i <= goal; i++ {
		for j := 1; j <= min(i, n); j++ {
			dp[i][j] = dp[i-1][j-1] * (n - j + 1) % 1_000_000_007
			if j > k {
				dp[i][j] = (dp[i][j] + dp[i-1][j]*(j-k)) % 1_000_000_007
			}
		}
	}

	return dp[goal][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
