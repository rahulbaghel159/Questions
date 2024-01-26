package graphgeneral

// https://leetcode.com/problems/clone-graph
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	return recursiveCloneGraph(node, make(map[*Node]*Node))
}

func recursiveCloneGraph(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if clone, exists := visited[node]; exists {
		return clone
	}

	clone := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	visited[node] = clone

	for i := 0; i < len(node.Neighbors); i++ {
		clone.Neighbors[i] = recursiveCloneGraph(node.Neighbors[i], visited)
	}

	return clone
}
