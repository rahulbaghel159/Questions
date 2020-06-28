package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/251/scenario-i-recurrence-relation/2378/

/*
The recursive version is slightly trickier and the key is to work backwards. Assume that the rest of the list had already been reversed, now how do I reverse the front part? Let's assume the list is: n1 → … → nk-1 → nk → nk+1 → … → nm → Ø
Assume from node nk+1 to nm had been reversed and you are at node nk.
n1 → … → nk-1 → nk → nk+1 ← … ← nm
We want nk+1’s next node to point to nk.
So,
nk.next.next = nk;
Be very careful that n1's next must point to Ø. If you forget about this, your linked list has a cycle in it. This bug could be caught if you test your code with a linked list of size 2.
*/

func reverseList(head *ListNode) *ListNode {
	//breaking condition
	if head == nil || head.Next == nil {
		return head
	}

	//recursing
	n := reverseList(head.Next)

	//processing
	head.Next.Next = head
	head.Next = nil
	return n
}
