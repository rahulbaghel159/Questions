package combinations

//https://leetcode.com/problems/combinations/

func combine(n int, k int) [][]int {
	//approach1
	//return recurse(n, k)
	//approach2
	return backtrack(n, 1, k, make([]int, 0), make([][]int, 0))
}

func recurse(n int, k int) [][]int {
	if k == 1 {
		result := make([][]int, n)
		for i := 1; i <= n; i++ {
			result[i-1] = make([]int, 1)
			result[i-1] = []int{i}
		}
		return result
	}

	result, newResult := recurse(n, k-1), make([][]int, 0)

	for i := 0; i < len(result); i++ {
		for j := 1; j <= n; j++ {
			if isValid(result[i], j) {
				s := make([]int, len(result[i])+1)

				for k := 0; k < len(s)-1; k++ {
					s[k] = result[i][k]
				}
				s[len(s)-1] = j

				newResult = append(newResult, s)
			}
		}
	}

	return newResult
}

func isValid(arr []int, n int) bool {
	unique, m := make([]int, 20), 0

	for _, v := range arr {
		if unique[v-1] > 0 || v <= m {
			return false
		}
		m = max(v, m)
		unique[v-1]++
	}

	if unique[n-1] > 0 || n <= m {
		return false
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func backtrack(n, firstNum, k int, curr []int, ans [][]int) [][]int {
	if len(curr) == k {
		ans = append(ans, curr)
		return ans
	}

	need, remain := k-len(curr), n-firstNum+1
	available := remain - need

	for i := firstNum; i <= firstNum+available; i++ {
		curr = custom_append(curr, i)
		ans = backtrack(n, i+1, k, curr, ans)
		curr = custom_slice(curr, len(curr)-1)
	}
	return ans
}

func custom_append(arr []int, n int) []int {
	res := make([]int, len(arr)+1)

	for i, v := range arr {
		res[i] = v
	}

	res[len(arr)] = n

	return res
}

func custom_slice(arr []int, idx int) []int {
	if len(arr) <= 1 {
		return []int{}
	}
	res := make([]int, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		res[i] = arr[i]
	}

	return res
}
