package linkedlist

// https://leetcode.com/explore/interview/card/google/60/linked-list-5/3064/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	first, second := dummy, dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
