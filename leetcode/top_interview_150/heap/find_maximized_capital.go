package heap

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/problems/ipo/
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]Project, n)

	for i := 0; i < n; i++ {
		projects[i] = Project{capital: capital[i], profit: profits[i]}
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	q, ptr := &MaxHeap{}, 0
	heap.Init(q)
	for i := 0; i < k; i++ {
		for ptr < n && projects[ptr].capital <= w {
			heap.Push(q, projects[ptr].profit)
			ptr++
		}
		if q.Len() == 0 {
			break
		}
		w += heap.Pop(q).(int)
	}

	return w
}

type Project struct {
	capital, profit int
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
