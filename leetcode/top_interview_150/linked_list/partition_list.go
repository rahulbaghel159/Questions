package linkedlist

// https://leetcode.com/problems/partition-list
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ptr, list1, list2 := head, &ListNode{Val: -999}, &ListNode{Val: -999}

	tmp1, tmp2 := list1, list2
	for ptr != nil {
		if ptr.Val < x {
			tmp1.Next = ptr
			tmp1 = tmp1.Next
		} else {
			tmp2.Next = ptr
			tmp2 = tmp2.Next
		}
		ptr = ptr.Next
	}

	tmp1.Next = list2.Next
	tmp2.Next = nil

	return list1.Next
}
