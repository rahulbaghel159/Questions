package main

import "sort"

// https://leetcode.com/explore/learn/card/fun-with-arrays/523/searching-for-items-in-an-array/3228/
func heightChecker(heights []int) int {
	if len(heights) <= 1 {
		return len(heights)
	}

	original := make([]int, len(heights))

	copy(original, heights)

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})

	res := 0

	for i := range heights {
		if heights[i] != original[i] {
			res++
		}
	}

	return res
}
