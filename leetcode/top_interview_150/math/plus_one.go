package math

// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	n, carry := len(digits), 0
	digits[n-1]++

	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] < 10 {
			carry = 0
			break
		}
		carry = digits[i] / 10
		digits[i] = digits[i] % 10
	}

	if carry > 0 {
		res := []int{carry}
		return append(res, digits...)
	}

	return digits
}
