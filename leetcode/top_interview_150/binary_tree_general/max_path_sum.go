package binarytreegeneral

// https://leetcode.com/problems/binary-tree-maximum-path-sum
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, maxSum := recursiveMaxPath(root, -9999999999999999)

	return maxSum
}

func recursiveMaxPath(node *TreeNode, maxSum int) (int, int) {
	if node == nil {
		return 0, maxSum
	}

	left_gain, maxSum := recursiveMaxPath(node.Left, maxSum)
	left_gain = max(left_gain, 0)

	right_gain, maxSum := recursiveMaxPath(node.Right, maxSum)
	right_gain = max(right_gain, 0)

	maxSum = max(maxSum, left_gain+right_gain+node.Val)

	return max(left_gain+node.Val, right_gain+node.Val), maxSum
}
