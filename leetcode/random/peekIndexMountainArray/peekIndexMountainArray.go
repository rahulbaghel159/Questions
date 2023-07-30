package peekIndexMountainArray

//https://leetcode.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	if len(arr) < 3 {
		return -1
	}

	start, end, index := 0, len(arr)-1, 0

	for start < end {
		mid := (end + start) / 2

		//peak found
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			index = mid
			break
		}

		if arr[mid] > arr[mid-1] && arr[mid] < arr[mid]+1 {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return index
}
