package multidp

// https://leetcode.com/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			dp[i][j] = 999999999999999999
		}
	}

	dp[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j+1 < len(dp[i]) {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+grid[i][j+1])
			}
			if i+1 < len(dp) {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+grid[i+1][j])
			}
		}
	}

	return dp[len(grid)-1][len(grid[len(grid)-1])-1]
}
