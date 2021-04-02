package maxarea

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, []int{0, h}...)
	verticalCuts = append(verticalCuts, []int{0, w}...)

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxh, maxw := 0, 0

	for i := 1; i < len(horizontalCuts); i++ {
		maxh = max(maxh, horizontalCuts[i]-horizontalCuts[i-1])
	}

	for i := 1; i < len(verticalCuts); i++ {
		maxw = max(maxw, verticalCuts[i]-verticalCuts[i-1])
	}

	return (maxh * maxw) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
