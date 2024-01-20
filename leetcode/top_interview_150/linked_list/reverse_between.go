package linkedlist

// https://leetcode.com/problems/reverse-linked-list-ii
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	if left == right {
		return head
	}

	var (
		curr, prev, con, tail, third *ListNode
	)
	curr, prev = head, nil

	for left > 1 {
		prev = curr
		curr = curr.Next
		left--
		right--
	}

	con, tail, third = prev, curr, nil
	for right > 0 {
		third = curr.Next
		curr.Next = prev
		prev = curr
		curr = third
		right--
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}

	tail.Next = curr

	return head
}
