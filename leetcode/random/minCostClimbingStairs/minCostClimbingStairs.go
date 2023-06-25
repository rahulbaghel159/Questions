package mincostclimbingstairs

import (
	"fmt"
	"strings"
)

func minCostClimbingStairs(cost []int) int {
	memo := make(map[int]int)
	return minCostClimbingStairsMemoized(cost, len(cost), memo)
}

func minCostClimbingStairsMemoized(cost []int, nStairs int, memo map[int]int) int {
	//str := stringify(cost)
	if v, ok := memo[nStairs]; ok {
		return v
	}

	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	if len(cost) <= 1 {
		return 0
	}

	n := len(cost)
	result := min(cost[n-1]+minCostClimbingStairsMemoized(copyArr(cost, 0, n-2), n-1, memo), cost[n-2]+minCostClimbingStairsMemoized(copyArr(cost, 0, n-3), n-2, memo))
	memo[nStairs] = result
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func copyArr(arr []int, start, end int) []int {
	result := make([]int, end-start+1)

	index := 0
	for i := start; i <= end; i++ {
		result[index] = arr[i]
		index++
	}

	return result
}

func stringify(arr []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(arr), " "), ","), "[]")
}
