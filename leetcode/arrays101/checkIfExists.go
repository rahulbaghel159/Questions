package main

import "sort"

// https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
func checkIfExist(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		target := 2 * arr[i]
		found, index := binarySearch(arr, target)
		if found && i != index {
			return true
		}
	}

	return false
}

func binarySearch(arr []int, num int) (bool, int) {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == num {
			return true, mid
		} else if arr[mid] > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false, -1
}
