package minoperationsarrayempty

import "math"

// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
// func minOperations(nums []int) int {
// 	count := make(map[int]int)

// 	for _, num := range nums {
// 		count[num]++
// 	}

// 	ans, memo := 0, make(map[int]int)
// 	for _, v := range count {
// 		tmp := countOperations(v, memo)
// 		if tmp == -1 {
// 			return -1
// 		}
// 		ans += tmp
// 	}

// 	return ans
// }

// func countOperations(num int, memo map[int]int) int {
// 	if _, ok := memo[num]; ok {
// 		return memo[num]
// 	}

// 	if num == 0 {
// 		return 0
// 	}

// 	if num < 0 {
// 		return -1
// 	}

// 	if num-2 == 0 || num-3 == 0 {
// 		return 1
// 	}

// 	left, right := countOperations(num-2, memo), countOperations(num-3, memo)

// 	if left == right && left == -1 {
// 		return -1
// 	}

// 	if max(left, 0) == 0 {
// 		memo[num] = max(right, 0) + 1
// 	} else if max(right, 0) == 0 {
// 		memo[num] = max(left, 0) + 1
// 	} else {
// 		memo[num] = min(max(left, 0), max(right, 0)) + 1
// 	}

// 	return memo[num]
// }

func minOperations(nums []int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	ans := 0
	for _, v := range count {
		if v == 1 {
			return -1
		}
		ans += int(math.Ceil(float64(v) / 3))
	}

	return ans
}
