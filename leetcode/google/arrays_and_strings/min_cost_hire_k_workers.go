package google

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3061/
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([][2]float64, len(quality))

	for i := 0; i < len(quality); i++ {
		workers[i] = [2]float64{float64(wage[i]) / float64(quality[i]), float64(quality[i])}
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i][0] < workers[j][0]
	})

	var res, qSum float64
	res, qSum = 999999999999999999, 0

	pq := PriorityQueue{}
	heap.Init(&pq)

	for _, worker := range workers {
		qSum += worker[1]
		heap.Push(&pq, &Item{
			priority: -worker[1],
		})
		if pq.Len() > k {
			qSum += heap.Pop(&pq).(*Item).priority
		}
		if pq.Len() == k {
			res = min(res, qSum*worker[0])
		}
	}

	return res
}

type Item struct {
	value    int
	priority float64
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:len(old)-1]
	return item
}
