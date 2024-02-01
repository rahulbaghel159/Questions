package divideandconquer

// https://leetcode.com/problems/sort-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := splitMid(head)
	left := sortList(head)
	right := sortList(mid)

	return merge(left, right)
}

func merge(node1, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			tmp.Next = node1
			node1 = node1.Next
		} else {
			tmp.Next = node2
			node2 = node2.Next
		}
		tmp = tmp.Next
	}

	if node1 != nil {
		tmp.Next = node1
	} else {
		tmp.Next = node2
	}

	return dummy.Next
}

func splitMid(head *ListNode) *ListNode {
	var midPrev *ListNode

	for head != nil && head.Next != nil {
		if midPrev == nil {
			midPrev = head
		} else {
			midPrev = midPrev.Next
		}

		head = head.Next.Next
	}

	mid := midPrev.Next
	midPrev.Next = nil

	return mid
}
