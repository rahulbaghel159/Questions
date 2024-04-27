package atlassian

// https://leetcode.com/problems/cherry-pickup/
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	if grid[n-1][n-1] == 1 {
		dp[n-1][n-1] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != -1 {
				right, down := 0, 0
				if i+1 < n {
					down = dp[i+1][j]
				}
				if j+1 < n {
					right = dp[i][j+1]
				}

				if right == 0 && down == 0 {
					continue
				}

				if down > right {
					dp[i][j] = grid[i][j] + down
					if i+1 < n {
						grid[i+1][j] = 0
					}
				} else {
					dp[i][j] = grid[i][j] + right
					if j+1 < n {
						grid[i][j+1] = 0
					}
				}
			}
		}
	}
	downTrip := dp[0][0]

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = 0
		}
	}
	if grid[n-1][n-1] == 1 {
		dp[n-1][n-1] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != -1 {
				right, down := 0, 0
				if i+1 < n {
					down = dp[i+1][j]
				}
				if j+1 < n {
					right = dp[i][j+1]
				}

				if right == 0 && down == 0 {
					continue
				}

				if down > right {
					dp[i][j] = grid[i][j] + down
					if i+1 < n {
						grid[i+1][j] = 0
					}
				} else {
					dp[i][j] = grid[i][j] + right
					if j+1 < n {
						grid[i][j+1] = 0
					}
				}
			}
		}
	}
	upTrip := dp[0][0]

	return downTrip + upTrip
}
