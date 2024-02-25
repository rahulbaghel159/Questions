package multidp

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	length := len(prices)
	leftMin, rightMax, leftProfit, rightProfit := prices[0], prices[length-1], make([]int, length), make([]int, length+1)

	for l := 1; l < length; l++ {
		leftProfit[l] = max(leftProfit[l-1], prices[l]-leftMin)
		leftMin = min(leftMin, prices[l])

		r := length - l - 1
		rightProfit[r] = max(rightProfit[r+1], rightMax-prices[r])
		rightMax = max(rightMax, prices[r])
	}

	res := 0
	for i := 0; i < length; i++ {
		res = max(res, leftProfit[i]+rightProfit[i+1])
	}

	return res
}
