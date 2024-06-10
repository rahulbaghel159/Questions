package dynamicprogramming

// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/304/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min_so_far, profit := prices[0], 0
	for _, price := range prices {
		profit = max(profit, price-min_so_far)
		min_so_far = min(min_so_far, price)
	}

	return profit
}
