package maximumsubsequencescore

import (
	"container/heap"
	"sort"
)

//https://leetcode.com/problems/maximum-subsequence-score/

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	if len(nums1) != len(nums2) {
		return int64(0)
	}

	var (
		best, total int64
	)

	h := &MinHeap{}

	items := make([]struct {
		val1 int
		val2 int
	}, len(nums1))

	for i := range items {
		items[i].val1 = nums1[i]
		items[i].val2 = nums2[i]
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].val2 > items[j].val2
	})

	heap.Init(h)

	for _, v := range items {
		heap.Push(h, v.val1)
		total += int64(v.val1)

		if h.Len() > k {
			t := heap.Pop(h)
			total -= int64(t.(int))
		}

		if h.Len() == k {
			best = max(best, total*int64(v.val2))
		}
	}

	return int64(best)
}

func max(x, y int64) int64 {
	if y > x {
		return y
	}
	return x
}
