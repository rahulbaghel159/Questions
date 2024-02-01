package kadanesalgo

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return recursiveMaxSubArray(nums, 0, len(nums)-1)
}

func recursiveMaxSubArray(nums []int, left, right int) int {
	if left > right {
		return -999999999999999999
	}

	mid := (left + right) / 2
	curr, left_best_sum, right_best_sum := 0, 0, 0

	for i := mid - 1; i >= left; i-- {
		curr += nums[i]
		left_best_sum = max(left_best_sum, curr)
	}

	curr = 0
	for i := mid + 1; i <= right; i++ {
		curr += nums[i]
		right_best_sum = max(right_best_sum, curr)
	}

	combined_sum := left_best_sum + nums[mid] + right_best_sum

	left_half := recursiveMaxSubArray(nums, left, mid-1)
	right_half := recursiveMaxSubArray(nums, mid+1, right)

	return max(combined_sum, max(left_half, right_half))
}

// func maxSubArray(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	res, tmp := -999999999999999999, 0

// 	for _, v := range nums {
// 		tmp += v
// 		res = max(res, tmp)
// 		if tmp < 0 {
// 			tmp = 0
// 		}
// 	}

// 	return res
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
