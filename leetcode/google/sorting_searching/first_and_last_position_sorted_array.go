package sortingsearching

// https://leetcode.com/explore/interview/card/google/63/sorting-and-searching-4/3081/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := findBounds(nums, target, true)
	if left == -1 {
		return []int{-1, -1}
	}

	right := findBounds(nums, target, false)

	return []int{left, right}
}

func findBounds(nums []int, target int, isStart bool) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			if isStart {
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
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
