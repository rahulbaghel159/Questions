package intervals

import "sort"

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	distinct, idx := make([][]int, 0), 0
	distinct = append(distinct, []int{points[0][0], points[0][1]})

	for i := 1; i < len(points); i++ {
		if isOverlapping(distinct[idx], points[i]) {
			distinct[idx] = []int{max(distinct[idx][0], points[i][0]), min(distinct[idx][1], points[i][1])}
		} else {
			distinct = append(distinct, points[i])
			idx++
		}
	}

	return len(distinct)
}

func isOverlapping(a, b []int) bool {
	return a[1] >= b[0]
}
