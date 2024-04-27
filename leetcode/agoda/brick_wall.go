package agoda

// https://leetcode.com/problems/brick-wall/
func leastBricks(wall [][]int) int {
	if len(wall) == 0 {
		return 0
	}

	edges_map := make(map[int]int)
	for _, bricks := range wall {
		length_so_far := 0
		for i := 0; i < len(bricks)-1; i++ {
			length_so_far += bricks[i]
			edges_map[length_so_far]++
		}
	}

	max_edges := 0
	for _, v := range edges_map {
		max_edges = max(max_edges, v)
	}

	return len(wall) - max_edges
}
