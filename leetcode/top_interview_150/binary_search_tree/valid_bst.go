package binarysearchtree

// https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return validateBST(root, nil, nil)
}

func validateBST(node, low, high *TreeNode) bool {
	if node == nil {
		return true
	}

	if (low != nil && node.Val <= low.Val) || (high != nil && node.Val >= high.Val) {
		return false
	}

	return validateBST(node.Right, node, high) && validateBST(node.Left, low, node)
}
