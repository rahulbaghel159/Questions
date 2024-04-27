package others

import (
	"math"
)

// https://leetcode.com/explore/interview/card/google/66/others-4/3096/
func reverse(x int) int {
	rev, is_negative := 0, false

	if x < 0 {
		is_negative = true
		x = -x
	}

	for x != 0 {
		rem := x % 10
		x = x / 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && rem > 7) {
			return 0
		}

		rev = rev*10 + rem
	}

	if is_negative {
		return -rev
	}

	return rev
}
