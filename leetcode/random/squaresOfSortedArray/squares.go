package squaresofsortedarray

// https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	left, right, i := 0, len(nums)-1, len(nums)-1

	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
		i--
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
