package atlassian

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-words/
func topKFrequent(words []string, k int) []string {
	count_map := make(map[string]int)

	for _, w := range words {
		count_map[w]++
	}

	word_heap := MinHeap{}
	heap.Init(&word_heap)
	for key, value := range count_map {
		if word_heap.Len() == k {
			curr := heap.Pop(&word_heap).(Node)
			if curr.value > value || (curr.value == value && curr.label < key) {
				heap.Push(&word_heap, curr)
			} else {
				heap.Push(&word_heap, Node{
					label: key,
					value: value,
				})
			}
		} else {
			heap.Push(&word_heap, Node{
				label: key,
				value: value,
			})
		}
	}

	ans, i := make([]string, k), k-1
	for word_heap.Len() > 0 {
		ans[i] = heap.Pop(&word_heap).(Node).label
		i--
	}

	return ans
}

type Node struct {
	label string
	value int
}

type MinHeap []Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].value == h[j].value {
		return h[i].label > h[j].label
	}

	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}
