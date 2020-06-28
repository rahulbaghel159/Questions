package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/255/recursion-memoization/1662/

var masterStairs = map[int]int{}

func climbStairs(n int) int {
	//breaking condition
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	//memoization and processing and recurse
	n1, ok := masterStairs[n-1]
	if !ok {
		n1 = climbStairs(n - 1)
		masterStairs[n-1] = n1
	}
	n2, ok := masterStairs[n-2]
	if !ok {
		n2 = climbStairs(n - 2)
		masterStairs[n-2] = n2
	}
	return n1 + n2
}
