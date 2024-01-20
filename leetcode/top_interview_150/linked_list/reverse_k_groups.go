package linkedlist

// https://leetcode.com/problems/reverse-nodes-in-k-group
// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	count, ptr := 0, head

// 	for count < k && ptr != nil {
// 		ptr = ptr.Next
// 		count++
// 	}

// 	if count == k {
// 		reversed_head := reverseList(head, k)

// 		head.Next = reverseKGroup(ptr, k)

// 		return reversed_head
// 	}

// 	return head
// }

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		ptr, ktail, new_head *ListNode
	)

	ptr, ktail, new_head = head, nil, nil

	for ptr != nil {
		count := 0
		ptr = head

		for count < k && ptr != nil {
			ptr = ptr.Next
			count++
		}

		if count == k {
			rev_head := reverseList(head, k)

			if new_head == nil {
				new_head = rev_head
			}

			if ktail != nil {
				ktail.Next = rev_head
			}

			ktail = head
			head = ptr
		}
	}

	if ktail != nil {
		ktail.Next = head
	}

	if new_head != nil {
		return new_head
	}
	return head
}

func reverseList(head *ListNode, k int) *ListNode {
	//assuming that the list has atleast k nodes
	var (
		new_head, ptr *ListNode
	)

	new_head, ptr = nil, head

	for k > 0 {
		next_node := ptr.Next

		ptr.Next = new_head
		new_head = ptr
		ptr = next_node

		k--
	}

	return new_head
}
