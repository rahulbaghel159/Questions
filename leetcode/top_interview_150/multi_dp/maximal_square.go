package multidp

// https://leetcode.com/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	rows, cols, max_so_far := len(matrix), len(matrix[0]), 0
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				a, b, c := 0, 0, 0
				if i-1 < rows && i-1 >= 0 && j-1 < cols && j-1 >= 0 {
					a = dp[i-1][j-1]
				}
				if j-1 < cols && j-1 >= 0 {
					b = dp[i][j-1]
				}
				if i-1 < rows && i-1 >= 0 {
					c = dp[i-1][j]
				}

				dp[i][j] = min(a, min(b, c)) + 1

				max_so_far = max(max_so_far, dp[i][j])
			}
		}
	}

	return max_so_far * max_so_far
}
