package pointers

// https://leetcode.com/problems/container-with-most-water
func maxArea(height []int) int {
	left, right, area := 0, len(height)-1, 0

	for left < right {
		area = max(area, (right-left)*min(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
