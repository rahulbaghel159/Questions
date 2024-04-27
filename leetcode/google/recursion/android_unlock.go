package recursion

// https://leetcode.com/explore/interview/card/google/62/recursion-4/484/
func numberOfPatterns(m int, n int) int {
	jump := map[int]map[int]int{
		1: {3: 2, 7: 4, 9: 5},
		2: {8: 5},
		3: {1: 2, 7: 5, 9: 6},
		4: {6: 5},
		6: {4: 5},
		7: {1: 4, 3: 5, 9: 8},
		8: {2: 5},
		9: {3: 6, 1: 5, 7: 8},
	}

	count := 4*dfsPattern(jump, map[int]struct{}{1: {}}, m, n, 1) + 4*dfsPattern(jump, map[int]struct{}{2: struct{}{}}, m, n, 2) + dfsPattern(jump, map[int]struct{}{5: struct{}{}}, m, n, 5)

	return count
}

func dfsPattern(jump map[int]map[int]int, used map[int]struct{}, m, n, curr int) int {
	count := 0

	if len(used) >= m {
		count++
	}

	if len(used) >= n {
		return count
	}

	for i := 1; i < 10; i++ {
		flag := false

		if _, ok := used[i]; !ok {
			if _, ok1 := jump[curr][i]; !ok1 {
				flag = true
			}
			if _, ok2 := used[jump[curr][i]]; ok2 {
				flag = true
			}

			if flag {
				used[i] = struct{}{}
				count += dfsPattern(jump, used, m, n, i)
				delete(used, i)
			}
		}
	}

	return count
}
