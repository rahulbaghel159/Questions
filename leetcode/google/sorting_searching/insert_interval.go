package sortingsearching

// https://leetcode.com/explore/interview/card/google/63/sorting-and-searching-4/445/
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	left, right, pos := 0, len(intervals)-1, len(intervals)
	for left <= right {
		mid := left + (right-left)/2

		if intervals[mid][0] > newInterval[0] {
			pos = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if pos != len(intervals) {
		intervals = append(intervals[:pos+1], intervals[pos:]...)
		intervals[pos] = newInterval
	} else {
		intervals = append(intervals, newInterval)
	}

	final_interval := make([][]int, 0)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] <= intervals[i][1] {
			intervals[i+1][0] = min(intervals[i][0], intervals[i+1][0])
			intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])

			intervals[i] = []int{-1, -1}
		}
	}

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] != -1 && intervals[i][1] != -1 {
			final_interval = append(final_interval, intervals[i])
		}
	}

	return final_interval
}
