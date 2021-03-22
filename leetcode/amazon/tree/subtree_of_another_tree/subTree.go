package subtree

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return traverse(s, t)
}

func traverse(s *TreeNode, t *TreeNode) bool {
	return s != nil && (isEqual(s, t) || traverse(s.Left, t) || traverse(s.Right, t))
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}

	return (s.Val == t.Val) && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}
