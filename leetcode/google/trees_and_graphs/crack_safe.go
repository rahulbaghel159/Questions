package treesandgraphs

import "math"

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3075/discuss/110264/Easy-DFS
func crackSafe(n int, k int) string {
	total, sb := int(math.Pow(float64(k), float64(n))), make([]byte, 0)
	for i := 0; i < n; i++ {
		sb = append(sb, '0')
	}

	visited := make(map[string]struct{})
	visited[string(sb)] = struct{}{}

	dfsCrack(&sb, visited, n, k, total)

	return string(sb)
}

func dfsCrack(sb *[]byte, visited map[string]struct{}, n, k, goal int) bool {
	if len(visited) == goal {
		return true
	}

	prev := (*sb)[len(*sb)-n+1:]
	for i := 0; i < k; i++ {
		next := append(prev, byte(i+'0'))
		if _, ok := visited[string(next)]; !ok {
			visited[string(next)] = struct{}{}
			*sb = append(*sb, byte(i+'0'))
			if dfsCrack(sb, visited, n, k, goal) {
				return true
			} else {
				delete(visited, string(next))
				*sb = (*sb)[:len(*sb)-1]
			}
		}
	}

	return false
}
