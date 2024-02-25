package multidp

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
func maxProfit2(k int, prices []int) int {
	n := len(prices)

	if n <= 0 || k <= 0 {
		return 0
	}

	if k/2 > n {
		res := 0
		for i := 1; i < n; i++ {
			res += max(0, prices[i]-prices[i-1])
		}
		return res
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = []int{-1000000000, -1000000000}
		}
	}

	dp[0][0][0] = 0
	dp[0][1][1] = -prices[0]

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			if j > 0 {
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}

	res := 0
	for j := 0; j < k; j++ {
		res = max(res, dp[n-1][j][0])
	}

	return res
}
