package arrays

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice, maxProfit := prices[0], 0

	for _, v := range prices {
		minPrice = min(minPrice, v)
		maxProfit = max(maxProfit, v-minPrice)
	}

	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
