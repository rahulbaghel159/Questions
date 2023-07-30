package nonoverlappinginterval

import "sort"

//https://leetcode.com/problems/non-overlapping-intervals/description/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	k, ans := -999999999999999999, 0

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for _, v := range intervals {
		if v[0] >= k {
			k = v[1]
		} else {
			ans++
		}
	}

	return ans
}
