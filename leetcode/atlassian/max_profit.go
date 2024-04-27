package atlassian

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minSoFar, profit := prices[0], 0
	for i := 0; i < len(prices); i++ {
		profit = max(profit, prices[i]-minSoFar)
		if minSoFar > prices[i] {
			minSoFar = prices[i]
		}
	}

	return profit
}
