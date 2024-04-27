package atlassian

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	//no rotation
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[right]
}
