package multidp

// https://leetcode.com/problems/triangle/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := make([][]int, len(triangle))
	dp[0] = []int{triangle[0][0]}
	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(dp[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+triangle[i][j], dp[i-1][j]+triangle[i][j])
			}
		}
	}

	n := len(dp)
	res := dp[n-1][0]
	for _, v := range dp[n-1] {
		res = min(res, v)
	}

	return res
}
