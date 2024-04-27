package google

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3059/
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return len(intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}

		return intervals[i][1] > intervals[j][1]
	})

	fmt.Println(intervals)

	ans, min_heap := 1, MinHeap{}
	heap.Init(&min_heap)

	heap.Push(&min_heap, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if len(min_heap) > 0 {
			end_time := HeapPeek(min_heap)
			if end_time <= intervals[i][0] {
				heap.Pop(&min_heap)
			}
		}
		heap.Push(&min_heap, intervals[i][1])
		ans = max(ans, len(min_heap))
	}

	return ans
}

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
	*h = old[:n-1]
	return x
}

func HeapPeek(h MinHeap) int {
	n := heap.Pop(&h).(int)
	heap.Push(&h, n)
	return n
}
