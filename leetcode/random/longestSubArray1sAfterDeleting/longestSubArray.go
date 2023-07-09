package longestsubarray1safterdeleting

//https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

func longestSubarray(nums []int) int {
	zeroCount, longestWindow, start := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}
		longestWindow = max(longestWindow, i-start)
	}

	return longestWindow
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
