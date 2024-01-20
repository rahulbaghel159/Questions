package intervals

// https://leetcode.com/problems/insert-interval
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = insertInterval(intervals, newInterval)

	ans := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		currInterval := []int{intervals[i][0], intervals[i][1]}
		for i < len(intervals) && doesIntervalOverlap(currInterval, intervals[i]) {
			currInterval = mergeIntervals(currInterval, intervals[i])
			i++
		}
		i--
		ans = append(ans, currInterval)
	}

	return ans
}

// returns true if intervals a and b have common element
func doesIntervalOverlap(a, b []int) bool {
	return min(a[1], b[1])-max(a[0], b[0]) >= 0
}

// return the interval having all the elements of intervals a and b.
func mergeIntervals(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func upperBound(intervals [][]int, newInterval []int) int {
	if len(intervals) == 0 {
		return 0
	}

	start, end, ans := 0, len(intervals)-1, len(intervals)

	for start <= end {
		mid := (start + end) / 2
		if intervals[mid][0] > newInterval[0] {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	index := upperBound(intervals, newInterval)

	if index != len(intervals) {
		intervals = append(intervals[:index+1], intervals[index:]...)
		intervals[index] = newInterval
	} else {
		intervals = append(intervals, newInterval)
	}

	return intervals
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
