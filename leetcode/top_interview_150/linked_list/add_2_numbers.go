package linkedlist

// https://leetcode.com/problems/add-two-numbers
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil && l2 == nil {
// 		return nil
// 	}

// 	if l1 == nil {
// 		return l2
// 	}

// 	if l2 == nil {
// 		return l1
// 	}

// 	l, carry := &ListNode{}, 0
// 	head, prev := l, l
// 	for l1 != nil || l2 != nil {
// 		a, b := 0, 0
// 		if l1 != nil {
// 			a = l1.Val
// 		}
// 		if l2 != nil {
// 			b = l2.Val
// 		}
// 		sum := a + b + carry
// 		l.Val = sum % 10
// 		carry = sum / 10

// 		prev = l
// 		l.Next = &ListNode{}
// 		l = l.Next
// 		if l1 != nil {
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			l2 = l2.Next
// 		}
// 	}

// 	if carry > 0 {
// 		l.Val = carry
// 		l.Next = nil
// 	} else {
// 		prev.Next = nil
// 	}

// 	return head
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil {
		return l2
	}

	head_l1, head_l2 := l1, l2

	for l1.Next != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1.Next != nil {
		for l1.Next != nil {
			l2.Next = &ListNode{Val: 0, Next: nil}
			l1 = l1.Next
			l2 = l2.Next
		}
	} else {
		for l2.Next != nil {
			l1.Next = &ListNode{Val: 0, Next: nil}
			l2 = l2.Next
			l1 = l1.Next
		}
	}

	carry := 0
	l1, l2 = head_l1, head_l2
	for l1.Next != nil && l2.Next != nil {
		sum := l1.Val + l2.Val + carry

		l1.Val = sum % 10
		carry = sum / 10

		l1 = l1.Next
		l2 = l2.Next
	}

	sum := l1.Val + l2.Val + carry

	l1.Val = sum % 10
	carry = sum / 10

	if carry > 0 {
		l1.Next = &ListNode{Val: carry, Next: nil}
	}

	return head_l1
}
