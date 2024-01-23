package binarytreegeneral

// https://leetcode.com/problems/path-sum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return recursivePathSum(root, targetSum, 0)
}

func recursivePathSum(node *TreeNode, targetSum int, sumSoFar int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return sumSoFar+node.Val == targetSum
	}

	return recursivePathSum(node.Left, targetSum, sumSoFar+node.Val) || recursivePathSum(node.Right, targetSum, sumSoFar+node.Val)
}
