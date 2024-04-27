package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.com/explore/interview/card/google/60/linked-list-5/3066/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	node_map, dummy := make(map[*Node]*Node), &Node{}
	tmp, tmp_head := dummy, head
	for tmp_head != nil {
		tmp.Next = &Node{
			Val: tmp_head.Val,
		}
		tmp = tmp.Next
		node_map[tmp_head] = tmp
		tmp_head = tmp_head.Next
	}

	tmp_head, tmp = head, dummy.Next
	for tmp_head != nil {
		random := tmp_head.Random
		tmp.Random = node_map[random]

		tmp = tmp.Next
		tmp_head = tmp_head.Next
	}

	return dummy.Next
}
