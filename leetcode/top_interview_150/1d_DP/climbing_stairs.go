package ddp

// https://leetcode.com/problems/climbing-stairs/
// var memo map[int]int

// func climbStairs(n int) int {
// 	memo = make(map[int]int)
// 	return climbStairsHelper(n, memo)
// }

// func climbStairsHelper(n int, memo map[int]int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	if n == 2 {
// 		return 2
// 	}

// 	if v, ok := memo[n]; ok {
// 		return v
// 	}

// 	memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)

// 	return memo[n]
// }

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
