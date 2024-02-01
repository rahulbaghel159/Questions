package kadanesalgo

// https://leetcode.com/problems/maximum-sum-circular-subarray/
func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, currMax, minSum, currMin, totalSum := nums[0], 0, nums[0], 0, 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]

		currMax = max(currMax, 0) + nums[i]
		maxSum = max(maxSum, currMax)

		currMin = min(currMin, 0) + nums[i]
		minSum = min(minSum, currMin)
	}

	if totalSum == minSum {
		return maxSum
	}

	return max(maxSum, totalSum-minSum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func maxSubarraySumCircular(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	n := len(nums)
// 	rightMax, suffixSum := make([]int, n), nums[n-1]
// 	rightMax[n-1] = nums[n-1]

// 	for i := n - 2; i >= 0; i-- {
// 		suffixSum += nums[i]
// 		rightMax[i] = max(rightMax[i+1], suffixSum)
// 	}

// 	maxSum, prefixSum, specialSum, currMax := nums[0], 0, nums[0], 0
// 	for i := 0; i < n; i++ {
// 		currMax = max(currMax, 0) + nums[i]
// 		maxSum = max(maxSum, currMax)

// 		prefixSum += nums[i]
// 		if i+1 < n {
// 			specialSum = max(specialSum, prefixSum+rightMax[i+1])
// 		}
// 	}

// 	return max(specialSum, maxSum)
// }
