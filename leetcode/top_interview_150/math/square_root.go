package math

// https://leetcode.com/problems/sqrtx/
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right, num := 2, x/2, 0
	for left <= right {
		mid := left + (right-left)/2
		num = mid * mid

		if num == x {
			return mid
		}

		if num > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
