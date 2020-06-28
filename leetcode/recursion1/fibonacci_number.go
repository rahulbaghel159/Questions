package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/255/recursion-memoization/1661/

var masterFib = map[int]int{}

func fib(n int) int {
	//breaking condition
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	//memoization
	var num int
	num, ok := masterFib[n]
	if !ok {
		//recursing and processing
		num = fib(n-1) + fib(n-2)
		masterFib[n] = num
	}
	return num
}
