package math

// https://leetcode.com/problems/factorial-trailing-zeroes/
func trailingZeroes(n int) int {
	zeroCount := 0

	for n > 0 {
		n /= 5
		zeroCount += n
	}

	return zeroCount
}
