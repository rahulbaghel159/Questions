package intervals

import "sort"

// https://leetcode.com/problems/merge-intervals
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	for _, v := range intervals {
		if len(v) != 2 {
			return [][]int{}
		}
	}

	//sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res, prev, idx := make([][]int, 0), intervals[0][1], 0
	res = append(res, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		if prev >= intervals[i][0] {
			m := max(intervals[i][1], res[idx][1])
			res[idx][1] = m
			prev = m
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
			idx++
			prev = intervals[i][1]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
