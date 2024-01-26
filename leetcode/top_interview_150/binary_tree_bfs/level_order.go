package binarytreebfsgo

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := arrayqueue.New()
	queue.Enqueue(root)

	res := make([][]int, 0)
	for queue.Size() > 0 {
		size := queue.Size()
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			n, _ := queue.Dequeue()
			node := n.(*TreeNode)

			tmp = append(tmp, node.Val)

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
