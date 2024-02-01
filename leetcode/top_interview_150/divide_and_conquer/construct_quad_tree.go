package divideandconquer

// https://leetcode.com/problems/construct-quad-tree/
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}

	node, n := &Node{}, len(grid)

	if hasSameValue(grid) {
		node.IsLeaf = true
		if grid[0][0] == 0 {
			node.Val = false
		} else {
			node.Val = true
		}
		return node
	} else {
		node.IsLeaf = false
		node.Val = false

		midRow, midCol := n/2, n/2

		// Initialize slices for the four sections
		grid1 := make([][]int, midRow)
		grid2 := make([][]int, midRow)
		grid3 := make([][]int, midRow)
		grid4 := make([][]int, midRow)

		for i := 0; i < midRow; i++ {
			grid1[i] = grid[i][:midCol]
			grid2[i] = grid[i][midCol:]
			grid3[i] = grid[midRow+i][:midCol]
			grid4[i] = grid[midRow+i][midCol:]
		}

		node.TopLeft = construct(grid1)
		node.TopRight = construct(grid2)
		node.BottomLeft = construct(grid3)
		node.BottomRight = construct(grid4)
	}

	return node
}

func hasSameValue(grid [][]int) bool {
	if len(grid) == 0 {
		return true
	}

	n, val := len(grid), grid[0][0]

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != val {
				return false
			}
		}
	}

	return true
}
