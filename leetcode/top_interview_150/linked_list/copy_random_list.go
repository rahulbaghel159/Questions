package linkedlist

// https://leetcode.com/problems/copy-list-with-random-pointer
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldNode, visited := head, make(map[*Node]*Node)
	newNode := &Node{Val: oldNode.Val}

	visited[oldNode] = newNode

	for oldNode != nil {
		newNode.Random = getClonedNode(visited, oldNode.Random)
		newNode.Next = getClonedNode(visited, oldNode.Next)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return visited[head]
}

func getClonedNode(visited map[*Node]*Node, node *Node) *Node {
	if node != nil {
		if _, exists := visited[node]; !exists {
			visited[node] = &Node{Val: node.Val}
		}
		return visited[node]
	}

	return nil
}
