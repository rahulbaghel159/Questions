package binarytreebfsgo

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res, level, queue := make([][]int, 0), 0, arrayqueue.New()
	queue.Enqueue(root)

	for queue.Size() > 0 {
		size, tmp := queue.Size(), make([]int, 0)
		level++

		for i := 0; i < size; i++ {
			n, _ := queue.Dequeue()
			node := n.(*TreeNode)

			if !isEven(level) {
				tmp = append(tmp, node.Val)
			} else {
				tmp = append([]int{node.Val}, tmp...)
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}

	return false
}
