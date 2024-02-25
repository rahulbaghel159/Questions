package multidp

// https://leetcode.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[len(obstacleGrid)-1])

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	dp[len(dp)-1][len(dp[len(dp)-1])-2] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
