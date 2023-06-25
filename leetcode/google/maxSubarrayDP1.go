package google

import "math"

func maxSubArrayHelper(nums []int, start, end int) int {
	//base case
	if end == start {
		return nums[start]
	}

	mid := (start + end) / 2

	leftMax := maxSubArrayHelper(nums, start, mid)
	rightMax := maxSubArrayHelper(nums, mid+1, end)

	//mid crossing sum
	leftSum := math.MinInt32
	sum := 0
	for i := mid; i >= start; i-- {
		sum = sum + nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= end; i++ {
		sum = sum + nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return max(max(leftMax, rightMax), max(leftSum+rightSum, max(leftSum, rightSum)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArrayDP1(nums []int) int {
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}
