package hashmap

// https://leetcode.com/problems/longest-consecutive-sequence
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	num_set, longestStreak := make(map[int]struct{}), 0

	for _, v := range nums {
		num_set[v] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := num_set[num-1]; !ok {
			currentNum, currentStreak := num, 1

			_, ok := num_set[currentNum+1]
			for ok {
				currentNum++
				currentStreak++
				_, ok = num_set[currentNum+1]
			}

			longestStreak = max(longestStreak, currentStreak)
		}
	}

	return longestStreak
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
