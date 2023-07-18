package add2numbers2

import "fmt"

//https://leetcode.com/problems/add-two-numbers-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	result := &ListNode{}

	t1, t2, node, carry := l1, l2, result, 0

	for t1 != nil || t2 != nil {
		val := 0
		if t1 != nil {
			val += t1.Val
			t1 = t1.Next
		}
		if t2 != nil {
			val += t2.Val
			t2 = t2.Next
		}
		val += carry

		val, carry = val%10, val/10

		node.Val = val
		node.Next = &ListNode{}
		node = node.Next
	}

	if carry > 0 {
		node.Val = carry
		node.Next = nil
	}

	result = reverse(result)

	for result.Val == 0 && result.Next != nil {
		result = result.Next
	}

	return result
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil

	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func (l *ListNode) print() {
	if l == nil {
		return
	}

	head := l

	fmt.Print(head.Val)

	for head.Next != nil {
		head = head.Next
		fmt.Print(",", head.Val)
	}

	fmt.Println()
}
