package jan2024

func minimumOperationsToMakeEqual(x int, y int) int {
	return helper(x, y)
}

func helper(x, y int) int {
	if x <= y {
		return y - x
	}

	res := x - y
	res = min(res, helper(x/5, y)+1+x%5)
	res = min(res, helper(x/5+1, y)+1+5-x%5)
	res = min(res, helper(x/11, y)+1+x%11)
	res = min(res, helper(x/11+1, y)+1+11-x%11)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
