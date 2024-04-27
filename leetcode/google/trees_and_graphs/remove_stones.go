package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3076/
func removeStones(stones [][]int) int {
	if len(stones) == 0 {
		return 0
	}

	adj := make(map[[2]int][][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		adj[[2]int{stones[i][0], stones[i][1]}] = make([][]int, 0)
	}

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				adj[[2]int{stones[i][0], stones[i][1]}] = append(adj[[2]int{stones[i][0], stones[i][1]}], []int{stones[j][0], stones[j][1]})
				adj[[2]int{stones[j][0], stones[j][1]}] = append(adj[[2]int{stones[j][0], stones[j][1]}], []int{stones[i][0], stones[i][1]})
			}
		}
	}

	visited, count := make(map[[2]int]struct{}), 0
	for i := 0; i < len(stones); i++ {
		if _, ok := visited[[2]int{stones[i][0], stones[i][1]}]; !ok {
			count++
			dfsStones(adj, visited, stones[i][0], stones[i][1])
		}
	}

	return len(stones) - count
}

func dfsStones(adj map[[2]int][][]int, visited map[[2]int]struct{}, x, y int) {
	visited[[2]int{x, y}] = struct{}{}

	for _, neighbour := range adj[[2]int{x, y}] {
		if _, ok := visited[[2]int{neighbour[0], neighbour[1]}]; !ok {
			dfsStones(adj, visited, neighbour[0], neighbour[1])
		}
	}
}
