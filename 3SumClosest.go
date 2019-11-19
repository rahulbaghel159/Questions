package threesum

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			current := nums[i] + nums[left] + nums[right]
			if current == target {
				return target
			}
			if current > target {
				right--
			}
			if current < target {
				left++
			}
			if abs(current-target) < abs(result-target) {
				result = current
			}
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}
