package arrays

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii

func maxProfit2(prices []int) int {
	peak, valley, i, maxProfit := prices[0], prices[0], 0, 0

	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]

		maxProfit += peak - valley
	}

	return maxProfit
}
