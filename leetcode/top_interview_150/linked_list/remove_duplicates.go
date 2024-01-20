package linkedlist

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: -999, Next: head}
	prev, ptr := dummy, dummy.Next

	for ptr != nil && ptr.Next != nil {
		found := false
		if ptr.Val == ptr.Next.Val {
			found = true
			v := ptr.Val
			for ptr != nil && ptr.Val == v {
				ptr = ptr.Next
			}
		}

		if found {
			prev.Next = ptr
		} else {
			prev = prev.Next
			ptr = ptr.Next
		}
	}

	if ptr == nil {
		prev.Next = nil
	}

	return dummy.Next
}
