package binarytreegeneral

// https://leetcode.com/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := depth(root)

	if d == 0 {
		return 1
	}

	left, right, pivot := 1, pow(2, d)-1, 0
	for left <= right {
		pivot = left + (right-left)/2
		if exists(root, d, pivot) {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return pow(2, d) - 1 + left
}

// func recursiveCount(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}

// 	if node.Left == nil && node.Right == nil {
// 		return 1
// 	}

// 	return recursiveCount(node.Left) + recursiveCount(node.Right) + 1
// }

func depth(node *TreeNode) int {
	d := 0
	for node.Left != nil {
		d++
		node = node.Left
	}

	return d
}

func exists(node *TreeNode, d, idx int) bool {
	left, right, pivot := 0, pow(2, d)-1, 0
	for i := 0; i < d; i++ {
		pivot = left + (right-left)/2
		if idx <= pivot {
			node = node.Left
			right = pivot
		} else {
			node = node.Right
			left = pivot + 1
		}
	}

	return node != nil
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		res *= a
		b--
	}

	return res
}
