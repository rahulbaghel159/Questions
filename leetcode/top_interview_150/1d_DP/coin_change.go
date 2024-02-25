package ddp

// https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	max, dp := amount+1, make([]int, amount+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = max
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
