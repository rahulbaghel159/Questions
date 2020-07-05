package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/256/complexity-analysis/2380/

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n == 0 {
		return 1.0
	}
	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}
