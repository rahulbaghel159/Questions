package binarysearch

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findBound(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}

	last := findBound(nums, target, false)

	return []int{first, last}
}

func findBound(nums []int, target int, isFirst bool) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			if isFirst {
				if mid == left || nums[mid-1] != target {
					return mid
				}
				right = mid - 1
			} else {
				if mid == right || nums[mid+1] != target {
					return mid
				}
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
