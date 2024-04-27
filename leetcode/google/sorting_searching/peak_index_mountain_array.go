package sortingsearching

// https://leetcode.com/explore/interview/card/google/63/sorting-and-searching-4/3084/
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2

		if mid == left && arr[mid+1] > arr[mid] {
			return mid + 1
		} else if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] && arr[mid+1] > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
