package binarytreegeneral

// https://leetcode.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return recursiveSumNumbers(root, 0, 0)
}

func recursiveSumNumbers(node *TreeNode, val, totalSum int) int {
	if node == nil {
		return totalSum
	}

	if node.Left == nil && node.Right == nil {
		totalSum += val*10 + node.Val
	}

	totalSum = recursiveSumNumbers(node.Left, val*10+node.Val, totalSum)
	totalSum = recursiveSumNumbers(node.Right, val*10+node.Val, totalSum)

	return totalSum
}
