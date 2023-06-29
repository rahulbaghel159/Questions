package minCostArrayEqual

import (
	"sort"
)

func minCost(nums []int, cost []int) int64 {
	if len(nums) == 0 || len(nums) != len(cost) {
		return 0
	}

	numsAndCost := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		numsAndCost[i] = make([]int, 2)

		numsAndCost[i][0] = nums[i]
		numsAndCost[i][1] = cost[i]
	}

	sort.Slice(numsAndCost, func(i, j int) bool {
		return numsAndCost[i][0] < numsAndCost[j][0]
	})

	prefixCost := make([]int64, len(nums))

	prefixCost[0] = int64(numsAndCost[0][1])

	for i := 1; i < len(cost); i++ {
		prefixCost[i] = prefixCost[i-1] + int64(numsAndCost[i][1])
	}

	minCost, prevCost := int64(0), int64(0)

	for i := 1; i < len(nums); i++ {
		minCost = minCost + int64((numsAndCost[i][0]-numsAndCost[0][0])*numsAndCost[i][1])
		prevCost = minCost
	}

	for i := 1; i < len(nums); i++ {
		gap := int64(numsAndCost[i][0] - numsAndCost[i-1][0])
		cost := prevCost + gap*prefixCost[i-1] - gap*(prefixCost[len(cost)-1]-prefixCost[i-1])
		if cost < minCost {
			minCost = cost
		}
		prevCost = cost
	}

	return int64(minCost)
}
