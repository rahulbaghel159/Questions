package atlassian

import "sort"

// https://leetcode.com/problems/set-intersection-size-at-least-two/
func intersectionSizeTwo(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	ans, a, b := 0, -1, -1
	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		if start > b {
			a, b = end-1, end
			ans += 2
		} else if start > a {
			a, b = b, end
			ans++
		}
	}

	return ans
}
