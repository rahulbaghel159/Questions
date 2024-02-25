package bitmanipulation

// https://leetcode.com/problems/bitwise-and-of-numbers-range/
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right = right & (right - 1)
	}

	return left & right
}
