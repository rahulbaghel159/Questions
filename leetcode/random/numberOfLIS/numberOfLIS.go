package numberoflis

//https://leetcode.com/problems/number-of-longest-increasing-subsequence/

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	arr, maxLen, count := make([][]int, 0), 0, 0

	for _, v := range nums {
		if len(arr) == 0 {
			arr[0] = []int{v}
			maxLen, count = 1, 1
			continue
		}
		newMax := maxLen
		for j := 0; j < len(arr); j++ {
			if isLIS(append(arr[j], v)) {
				arr[j] = append(arr[j], v)
				newMax = max(len(arr[j]), newMax)
			}
		}

		if newMax > maxLen {
			maxLen = newMax
			count = 1
		}
	}

	return count
}

func isLIS(arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	lis, prev := true, arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] <= prev {
			lis = false
			break
		}
		prev = arr[i]
	}

	return lis
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
