package google

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	//initially both min and max 0
	min, profit := prices[0], 0

	//if on moving right, price decreases, we can buy. If not, we can check the profit
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if profit < prices[i]-min {
				profit = prices[i] - min
			}
		}
	}

	return profit
}
