package binarytreegeneral

// https://leetcode.com/problems/binary-tree-maximum-path-sum
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left_sum := maxPathSum(root.Left)
	right_sum := maxPathSum(root.Right)

	return max(max(left_sum, right_sum), root.Val+left_sum+right_sum)
}
