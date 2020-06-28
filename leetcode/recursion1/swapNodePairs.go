package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/250/principle-of-recursion/1681/

func swapPairs(head *ListNode) *ListNode {
	//breaking condition
	if head == nil || head.Next == nil {
		return head
	}

	//processing
	l, r, n := head, head.Next, head.Next.Next
	r.Next = l
	l.Next = swapPairs(n)

	//recursing
	return r
}
