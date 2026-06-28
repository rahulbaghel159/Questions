package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/searching-for-items-in-an-array/3259/
func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		return []int{-1}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = findRightMax(arr, i)
	}

	return arr
}

func findRightMax(arr []int, index int) int {
	if len(arr) == 0 || index == len(arr)-1 {
		return -1
	}

	if len(arr) == 1 {
		return arr[0]
	}

	res := arr[index+1]

	for i := index + 1; i < len(arr); i++ {
		res = max(res, arr[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
