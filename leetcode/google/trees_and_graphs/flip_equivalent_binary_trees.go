package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3077/
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return recursiveFlipEquiv(root1, root2)
}

func recursiveFlipEquiv(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if (node1 == nil || node2 == nil) || (node1.Val != node2.Val) {
		return false
	}

	return (recursiveFlipEquiv(node1.Left, node2.Left) && recursiveFlipEquiv(node1.Right, node2.Right)) || (recursiveFlipEquiv(node1.Left, node2.Right) && recursiveFlipEquiv(node1.Right, node2.Left))
}
