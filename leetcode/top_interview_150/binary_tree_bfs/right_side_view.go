package binarytreebfsgo

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/binary-tree-right-side-view/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue, res := arrayqueue.New(), make([]int, 0)
	queue.Enqueue(root)

	for queue.Size() > 0 {
		size := queue.Size()
		for i := 0; i < size; i++ {
			n, _ := queue.Dequeue()
			node := n.(*TreeNode)

			if i == size-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}

	return res
}
