package twilio

// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
func minSwaps(s string) int {
	if len(s) == 0 {
		return 0
	}

	open, closed, res := 0, 0, 0

	for _, r := range s {
		if r == '[' {
			open++
		} else {
			closed++
		}

		if closed > open {
			closed--
			open++
			res++
		}
	}

	return res
}
