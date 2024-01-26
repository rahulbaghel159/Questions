package binarytreegeneral

import "github.com/emirpasic/gods/stacks/arraystack"

// https://leetcode.com/problems/binary-search-tree-iterator
type BSTIterator struct {
	stack *arraystack.Stack
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		stack: arraystack.New(),
	}

	iter.leftMostInOrder(root)

	return iter
}

func (b *BSTIterator) leftMostInOrder(node *TreeNode) {
	for node != nil {
		b.stack.Push(node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	n, _ := this.stack.Pop()
	node := n.(*TreeNode)

	if node.Right != nil {
		this.leftMostInOrder(node.Right)
	}

	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.stack.Size() > 0
}
