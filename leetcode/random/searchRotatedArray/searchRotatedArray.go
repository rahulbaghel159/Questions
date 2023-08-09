package searchrotatedarray

//https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	left, right, pivot := 0, len(nums)-1, 0

	for left <= right {
		mid := (right + left) / 2

		if mid > 0 && nums[mid-1] > nums[mid] {
			pivot = mid - 1
			break
		} else if mid < len(nums)-1 && nums[mid+1] < nums[mid] {
			pivot = mid
			break
		} else if nums[left] > nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	l := binarySearch(nums, 0, pivot-1, target)
	if l != -1 {
		return l
	}

	return binarySearch(nums, pivot, len(nums)-1, target)
}

func binarySearch(arr []int, leftBoundary, rightBoundary, target int) int {
	left, right := leftBoundary, rightBoundary

	for left <= right {
		mid := (left + right) / 2

		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
