package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/253/conclusion/2382/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//breaking condition
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	min := minListNode(l1, l2)
	max := maxListNode(l1, l2)
	if l1.Next == nil && l2.Next == nil {
		min.Next = max
		return min
	}

	//processing
	head := min
	for min.Next != nil && min.Next.Val < max.Val {
		min = min.Next
	}
	other := max.Next
	max.Next = min.Next
	min.Next = max

	//recursion
	return mergeTwoLists(head, other)
}

func minListNode(l1, l2 *ListNode) *ListNode {
	if l1.Val <= l2.Val {
		return l1
	}
	return l2
}

func maxListNode(l1, l2 *ListNode) *ListNode {
	if l1.Val > l2.Val {
		return l1
	}
	return l2
}
