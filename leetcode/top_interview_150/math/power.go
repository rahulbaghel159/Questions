package math

// https://leetcode.com/problems/powx-n/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	result := float64(1)

	for n != 0 {
		if n%2 == 1 {
			result *= x
			n -= 1
		}

		x *= x
		n /= 2
	}

	return result
}
