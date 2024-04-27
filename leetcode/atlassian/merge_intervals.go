package atlassian

import "sort"

// https://leetcode.com/problems/merge-intervals/
func mergeInterval(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i+1 < len(intervals); i++ {
		if intervals[i+1][0] <= intervals[i][1] {
			intervals[i+1][0] = min(intervals[i+1][0], intervals[i][0])
			intervals[i+1][1] = max(intervals[i+1][1], intervals[i][1])

			intervals[i][0], intervals[i][1] = -1, -1
		}
	}

	ans := make([][]int, 0)
	for _, interval := range intervals {
		if interval[0] != -1 && interval[1] != -1 {
			ans = append(ans, interval)
		}
	}

	return ans
}
