package binarytreegeneral

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := arrayqueue.New()
	queue.Enqueue(root)

	for queue.Size() > 0 {
		size := queue.Size()

		for i := 0; i < size; i++ {
			n, _ := queue.Dequeue()
			node := n.(*Node)

			if i < size-1 {
				p, _ := queue.Peek()
				node.Next = p.(*Node)
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}

	return root
}
