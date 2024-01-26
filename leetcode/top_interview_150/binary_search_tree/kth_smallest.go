package binarysearchtree

// https://leetcode.com/problems/kth-smallest-element-in-a-bst
func kthSmallest(root *TreeNode, k int) int {
	_, res := recursiveKthSmallest(root, k, 0)
	return res
}

func recursiveKthSmallest(node *TreeNode, k, curr int) (int, int) {
	if node == nil {
		return curr, -1
	}

	res := 0
	curr, res = recursiveKthSmallest(node.Left, k, curr)
	if curr == k {
		return curr, res
	}
	curr++
	if curr == k {
		return curr, node.Val
	}

	curr, res = recursiveKthSmallest(node.Right, k, curr)

	return curr, res
}
