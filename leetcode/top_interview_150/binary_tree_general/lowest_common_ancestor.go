package binarytreegeneral

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// refer to https://www.geeksforgeeks.org/lowest-common-ancestor-binary-tree-set-1/ for solution
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left_lca := lowestCommonAncestor(root.Left, p, q)
	right_lca := lowestCommonAncestor(root.Right, p, q)

	if left_lca != nil && right_lca != nil {
		return root
	}

	node := left_lca
	if left_lca == nil {
		node = right_lca
	}

	return node
}
