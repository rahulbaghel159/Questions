package recursion2

const (
	intMax = int(^uint(0) >> 1)
	intMin = -intMax - 1
)

func isValidBST(root *TreeNode) bool {
	return validateBST(root, intMin, intMax)
}

func validateBST(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return validateBST(root.Left, lower, root.Val) && validateBST(root.Right, root.Val, upper)
}
