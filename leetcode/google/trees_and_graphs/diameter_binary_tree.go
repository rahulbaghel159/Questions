package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3074/
var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter = 0
	recursiveDiameter(root)

	return diameter
}

func recursiveDiameter(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	left_gain := recursiveDiameter(node.Left)
	right_gain := recursiveDiameter(node.Right)

	diameter = max(diameter, left_gain+right_gain)
	gain := max(left_gain, right_gain) + 1

	return gain
}
