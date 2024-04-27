package dynamicprogramming

import "math"

// https://leetcode.com/explore/interview/card/google/64/dynamic-programming-4/3088/
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	return recursiveCoinChange(coins, amount, make(map[int]int, 0))
}

func recursiveCoinChange(coins []int, amount int, memo map[int]int) int {
	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	if _, ok := memo[amount]; ok {
		return memo[amount]
	}

	result := math.MaxInt
	for _, coin := range coins {
		temp_result := recursiveCoinChange(coins, amount-coin, memo)
		if temp_result != -1 {
			result = min(result, temp_result)
		}
	}

	if result == math.MaxInt {
		memo[amount] = -1
	} else {
		memo[amount] = 1 + result
	}

	return memo[amount]
}
