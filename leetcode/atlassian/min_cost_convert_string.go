package atlassian

// https://leetcode.com/problems/minimum-cost-to-convert-string-i/
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	if source == target {
		return 0
	}

	graph, INF := make([][]int, 26), 100000000
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 26)
		for j := 0; j < len(graph[i]); j++ {
			if i != j {
				graph[i][j] = INF
			} else {
				graph[i][j] = 0
			}
		}
	}

	for i := 0; i < len(original); i++ {
		graph[original[i]-'a'][changed[i]-'a'] = min(graph[original[i]-'a'][changed[i]-'a'], cost[i])
	}

	for k := 0; k < len(graph); k++ {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}

	var ans int64
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			if graph[source[i]-'a'][target[i]-'a'] == INF {
				return -1
			}
			ans += int64(graph[source[i]-'a'][target[i]-'a'])
		}
	}

	return ans
}
