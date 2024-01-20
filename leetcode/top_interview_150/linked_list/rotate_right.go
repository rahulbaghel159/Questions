package linkedlist

// https://leetcode.com/problems/rotate-list
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	old_tail, n := head, 1

	for old_tail.Next != nil {
		old_tail = old_tail.Next
		n++
	}
	old_tail.Next = head

	new_tail := head
	for i := 0; i < n-k%n-1; i++ {
		new_tail = new_tail.Next
	}
	new_head := new_tail.Next

	new_tail.Next = nil

	return new_head
}
