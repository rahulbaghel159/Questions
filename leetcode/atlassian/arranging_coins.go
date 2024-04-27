package atlassian

// https://leetcode.com/problems/arranging-coins/
// func arrangeCoins(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	res := 0
// 	for i := 1; n > 0; i++ {
// 		n -= i
// 		if n >= 0 {
// 			res++
// 		}
// 	}
// 	return res
// }

func arrangeCoins(n int) int {
	left, right, k, curr := 0, n, 0, 0
	for left < right {
		k = left + (right-left)/2
		curr = (k * (k + 1)) / 2
		if curr == n {
			return k
		} else if n < curr {
			right = curr - 1
		} else {
			left = curr + 1
		}
	}

	return right
}
