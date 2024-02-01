package divideandconquer

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var sortedList *ListNode

	for i := 0; i < len(lists); i++ {
		sortedList = mergeList(sortedList, lists[i])
	}

	return sortedList
}

func mergeList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}

	if list1 != nil {
		tmp.Next = list1
	} else {
		tmp.Next = list2
	}

	return dummy.Next
}
