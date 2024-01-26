package binarysearchtree

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	MinDiff  int
	PrevNode *TreeNode
)

func getMinimumDifference(root *TreeNode) int {
	MinDiff = 9999999
	PrevNode = nil
	inorderTraversal(root)
	return MinDiff
}

func inorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left)
	if PrevNode != nil {
		MinDiff = min(MinDiff, abs(node.Val-PrevNode.Val))
	}

	PrevNode = node
	inorderTraversal(node.Right)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
