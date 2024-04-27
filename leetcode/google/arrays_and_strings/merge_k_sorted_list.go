package google

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	INF := 99999999999999
	ans := &ListNode{
		Val: 0,
	}
	tmp_ans := ans

	for {
		min_index := -1
		tmp := &ListNode{
			Val: INF,
		}
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil && tmp.Val > lists[i].Val {
				tmp = lists[i]
				min_index = i
			}
		}

		if tmp.Val == INF {
			break
		}

		tmp_ans.Next = tmp
		tmp_ans = tmp_ans.Next
		lists[min_index] = lists[min_index].Next
	}

	return ans.Next
}
