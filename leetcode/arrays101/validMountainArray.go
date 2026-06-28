package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	topFound := false

	for i := 1; i < len(arr)-1; i++ {
		if !topFound {
			if arr[i] > arr[i-1] {
				if arr[i] > arr[i+1] {
					topFound = true
				}
			} else {
				return false
			}
		} else {
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	}

	return arr[len(arr)-2] > arr[len(arr)-1]
}
