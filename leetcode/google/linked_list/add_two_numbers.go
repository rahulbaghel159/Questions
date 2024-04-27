package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/explore/interview/card/google/60/linked-list-5/3063/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	tmp1, tmp2, carry, ans := l1, l2, 0, &ListNode{}
	tmp_ans := ans
	for tmp1 != nil && tmp2 != nil {
		v := tmp1.Val + tmp2.Val + carry
		if v > 9 {
			carry = v / 10
			v = v % 10
		} else {
			carry = 0
		}
		tmp_ans.Next = &ListNode{
			Val: v,
		}
		tmp_ans = tmp_ans.Next
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}

	for tmp1 != nil {
		v := tmp1.Val + carry
		if v > 9 {
			carry = v / 10
			v = v % 10
		} else {
			carry = 0
		}
		tmp_ans.Next = &ListNode{
			Val: v,
		}
		tmp_ans = tmp_ans.Next
		tmp1 = tmp1.Next
	}

	for tmp2 != nil {
		v := tmp2.Val + carry
		if v > 9 {
			carry = v / 10
			v = v % 10
		} else {
			carry = 0
		}
		tmp_ans.Next = &ListNode{
			Val: v,
		}
		tmp_ans = tmp_ans.Next
		tmp2 = tmp2.Next
	}

	if carry > 0 {
		tmp_ans.Next = &ListNode{
			Val: carry,
		}
	}

	return ans.Next
}
