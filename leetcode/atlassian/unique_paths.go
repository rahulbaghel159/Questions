package atlassian

// https://leetcode.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}

	dp[len(dp)-1][len(dp[0])-1] = 1
	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp[i]) - 1; j >= 0; j-- {
			down, right := 0, 0
			if i+1 < len(obstacleGrid) && obstacleGrid[i+1][j] != 1 {
				down = dp[i+1][j]
			}

			if j+1 < len(obstacleGrid[0]) && obstacleGrid[i][j+1] != 1 {
				right = dp[i][j+1]
			}
			dp[i][j] = max(dp[i][j], down+right)
		}
	}

	return dp[0][0]
}
