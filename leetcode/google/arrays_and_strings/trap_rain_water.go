package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/341/
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n, water := len(height), 0
	left_max, right_max := make([]int, n), make([]int, n)
	left_max[0], right_max[n-1] = height[0], height[n-1]

	for i := 0; i < n; i++ {
		if i > 0 {
			left_max[i] = max(height[i], left_max[i-1])
		}
		if i < n-1 {
			right_max[n-i-2] = max(height[n-i-2], right_max[n-i-1])
		}
	}

	for i := 0; i < n; i++ {
		water += (min(left_max[i], right_max[i]) - height[i])
	}

	return water
}
