package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3048/
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	maxArea, start, end := 0, 0, len(height)-1
	for start < end {
		area := min(height[start], height[end]) * (end - start)
		maxArea = max(maxArea, area)
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}

	return maxArea
}
