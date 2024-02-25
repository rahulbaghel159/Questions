package graphbfs

import "container/heap"

// https://leetcode.com/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}

	bank_map, allowed := make(map[string]bool, 0), []rune{'A', 'C', 'G', 'T'}

	for _, v := range bank {
		bank_map[v] = true
	}

	visited := make(map[string]bool, 0)

	pq := StrPriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &StrItem{
		value:    startGene,
		priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*StrItem)
		if item.value == endGene {
			return item.priority
		}
		if visited[item.value] {
			continue
		}
		str := []rune(item.value)
		for i := 0; i < len(str); i++ {
			ch := str[i]
			for _, r := range allowed {
				str[i] = r
				if bank_map[string(str)] {
					heap.Push(&pq, &StrItem{
						value:    string(str),
						priority: item.priority + 1,
					})
				}
			}
			str[i] = ch
		}
		visited[item.value] = true
	}

	return -1
}

type StrItem struct {
	value    string
	priority int
	index    int
}

type StrPriorityQueue []*StrItem

func (s StrPriorityQueue) Len() int           { return len(s) }
func (s StrPriorityQueue) Less(i, j int) bool { return s[i].priority < s[j].priority }
func (s StrPriorityQueue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].index = i
	s[j].index = j
}

func (s *StrPriorityQueue) Push(x any) {
	item := x.(*StrItem)
	*s = append(*s, item)
}

func (s *StrPriorityQueue) Pop() any {
	old := *s
	n := len(old)
	top := old[n-1]
	old[n-1] = nil
	top.index = -1
	*s = old[:n-1]
	return top
}
