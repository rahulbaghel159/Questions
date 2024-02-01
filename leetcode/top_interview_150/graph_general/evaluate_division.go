package graphgeneral

// https://leetcode.com/problems/evaluate-division
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	for i := 0; i < len(equations); i++ {
		dividend, divisor, quotient := equations[i][0], equations[i][1], values[i]

		if _, exists := graph[dividend]; !exists {
			graph[dividend] = make(map[string]float64)
		}

		if _, exists := graph[divisor]; !exists {
			graph[divisor] = make(map[string]float64)
		}

		graph[dividend][divisor] = quotient
		graph[divisor][dividend] = 1 / quotient
	}

	results := make([]float64, len(queries))

	for i := 0; i < len(queries); i++ {
		dividend, divisor := queries[i][0], queries[i][1]

		_, exists1 := graph[dividend]
		_, exists2 := graph[divisor]

		if !exists1 || !exists2 {
			results[i] = -1
		} else if dividend == divisor {
			results[i] = 1
		} else {
			visited := make(map[string]bool)
			results[i] = backtrackEvaluate(graph, dividend, divisor, 1, visited)
		}
	}

	return results
}

func backtrackEvaluate(graph map[string]map[string]float64, currNode, targetNode string, accProduct float64, visited map[string]bool) float64 {
	visited[currNode] = true
	ret := -1.0

	neighbors := graph[currNode]
	if _, exists := neighbors[targetNode]; exists {
		ret = accProduct * neighbors[targetNode]
	} else {
		for nextNode, value := range neighbors {
			if visited[nextNode] {
				continue
			}

			ret = backtrackEvaluate(graph, nextNode, targetNode, accProduct*value, visited)
			if ret != -1 {
				break
			}
		}
	}

	visited[currNode] = false
	return ret
}
