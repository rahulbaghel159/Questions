package backtracking

// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	return combineHelper(n, k)
}

func combineHelper(n, k int) [][]int {
	if k == 1 {
		res := make([][]int, n)
		for i := 1; i <= n; i++ {
			res[i-1] = []int{i}
		}
		return res
	}

	arr, res := combineHelper(n, k-1), make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		last := arr[i][len(arr[i])-1]
		for j := last + 1; j <= n; j++ {
			tmp := make([]int, len(arr[i])+1)
			for k := 0; k < len(arr[i]); k++ {
				tmp[k] = arr[i][k]
			}
			tmp[len(tmp)-1] = j
			res = append(res, tmp)
		}
	}

	return res
}
