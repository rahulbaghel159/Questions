package sortingsearching

import "sort"

// https://leetcode.com/explore/interview/card/google/63/sorting-and-searching-4/450/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] <= intervals[i][1] {
			intervals[i+1][0] = min(intervals[i][0], intervals[i+1][0])
			intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])

			intervals[i][0], intervals[i][1] = -1, -1
		}
	}

	result := make([][]int, 0)
	for _, interval := range intervals {
		if interval[0] != -1 && interval[1] != -1 {
			result = append(result, interval)
		}
	}

	return result
}
