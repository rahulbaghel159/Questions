package arrays

//https://leetcode.com/problems/trapping-rain-water

func trap(height []int) int {
	ans, left_max, right_max := 0, make([]int, len(height)), make([]int, len(height))

	for i := range height {
		left_max[i] = max(height[i], left_max[max(i-1, 0)])
	}

	for i := len(height) - 1; i >= 0; i-- {
		right_max[i] = max(right_max[min(i+1, len(height)-1)], height[i])
	}

	for i := range height {
		ans += min(left_max[i], right_max[i]) - height[i]
	}

	return ans
}
