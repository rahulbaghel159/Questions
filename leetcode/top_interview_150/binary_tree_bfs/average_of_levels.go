package binarytreebfsgo

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	queue := arrayqueue.New()
	queue.Enqueue(root)

	res := []float64{}
	for queue.Size() > 0 {
		size := queue.Size()

		sum, count := 0, 0
		for i := 0; i < size; i++ {
			n, _ := queue.Dequeue()
			node := n.(*TreeNode)

			sum += node.Val
			count++

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		avg := float64(sum) / float64(count)
		res = append(res, avg)
	}

	return res
}
