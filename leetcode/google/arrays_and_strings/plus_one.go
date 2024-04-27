package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/339/
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	digits[len(digits)-1] = digits[len(digits)-1] + 1
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + carry
		if digits[i] > 9 {
			carry = digits[i] / 10
			digits[i] = digits[i] % 10
		} else {
			carry = 0
		}
	}

	if carry > 0 {
		digits = append(digits, -1)
		for i := len(digits) - 1; i > 0; i-- {
			digits[i] = digits[i-1]
		}
		digits[0] = carry
	}

	return digits
}
