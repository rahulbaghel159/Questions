package heap

import (
	"container/heap"
	"strconv"
)

// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := PriorityQueue{}
	heap.Init(&pq)
	visited := make(map[string]struct{})

	heap.Push(&pq, &Item{
		value:    []int{0, 0},
		priority: nums1[0] + nums2[0],
	})
	visited["0_0"] = struct{}{}

	res := make([][]int, 0)

	for k > 0 && pq.Len() > 0 {
		k--
		item := heap.Pop(&pq).(*Item)
		i, j := item.value[0], item.value[1]
		res = append(res, []int{nums1[i], nums2[j]})

		if i+1 < len(nums1) {
			key := strconv.Itoa(i+1) + "_" + strconv.Itoa(j)
			if _, ok := visited[key]; !ok {
				visited[key] = struct{}{}
				heap.Push(&pq, &Item{
					value:    []int{i + 1, j},
					priority: nums1[i+1] + nums2[j],
				})

			}
		}

		if j+1 < len(nums2) {
			key := strconv.Itoa(i) + "_" + strconv.Itoa(j+1)
			if _, ok := visited[key]; !ok {
				visited[key] = struct{}{}
				heap.Push(&pq, &Item{
					value:    []int{i, j + 1},
					priority: nums1[i] + nums2[j+1],
				})
			}
		}
	}

	return res
}

type Item struct {
	value    []int
	priority int
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
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	top := old[n-1]
	top.index = -1
	old[n-1] = nil
	*pq = old[:n-1]
	return top
}
