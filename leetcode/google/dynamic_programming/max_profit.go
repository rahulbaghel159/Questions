package dynamicprogramming

// https://leetcode.com/explore/interview/card/google/64/dynamic-programming-4/3086/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max_profit, min_price_so_far := 0, prices[0]

	for _, price := range prices {
		min_price_so_far = min(min_price_so_far, price)

		max_profit = max(max_profit, price-min_price_so_far)
	}

	return max_profit
}
