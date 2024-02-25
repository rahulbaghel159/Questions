package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
