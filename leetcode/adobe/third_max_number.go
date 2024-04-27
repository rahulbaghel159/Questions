package adobe

import "math"

// https://leetcode.com/problems/third-maximum-number/
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	first, second, third := math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if num > first {
			third = second
			second = first
			first = num
		} else if num > second && num < first {
			third = second
			second = num
		} else if num > third && num < second {
			third = num
		}
	}

	if third == math.MinInt {
		return first
	}

	return third
}
