package linkedlist

// https://leetcode.com/problems/remove-nth-node-from-end-of-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	first, second := dummy, dummy

	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	if second != nil && second.Next != nil {
		second.Next = second.Next.Next
	}

	return dummy.Next
}
