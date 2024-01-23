package binarytreegeneral

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
func flatten(root *TreeNode) {
	if root != nil {
		curr := root
		for curr != nil {
			if curr.Left != nil {
				tmp, left, right_most := curr.Right, curr.Left, curr.Left
				for right_most.Right != nil {
					right_most = right_most.Right
				}
				curr.Right = left
				right_most.Right = tmp
				curr.Left = nil
				curr = left
			} else {
				curr = curr.Right
			}
		}
	}
}
