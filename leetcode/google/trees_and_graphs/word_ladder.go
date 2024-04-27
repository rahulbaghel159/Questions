package treesandgraphs

import "container/heap"

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3068/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	wordDict := make(map[string]bool)
	for _, word := range wordList {
		wordDict[word] = true
	}

	allowed := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
		'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	pq := StrPriorityQueue{}
	heap.Init(&pq)

	heap.Push(&pq, &strItem{
		value:    beginWord,
		priority: 0,
	})

	visited := map[string]int{
		beginWord: 0,
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*strItem)
		str, cost := node.value, node.priority
		if str == endWord {
			return cost + 1
		}

		if _, ok := visited[str]; ok && visited[str] < cost {
			continue
		}

		arr := []rune(str)
		for i := 0; i < len(arr); i++ {
			ch := arr[i]
			for _, r := range allowed {
				arr[i] = r

				if _, ok := visited[string(arr)]; (!ok || visited[string(arr)] > cost+1) && wordDict[string(arr)] {
					visited[string(arr)] = cost + 1
					heap.Push(&pq, &strItem{
						value:    string(arr),
						priority: cost + 1,
					})
				}
			}
			arr[i] = ch
		}
	}

	return 0
}

type strItem struct {
	value    string
	priority int
	index    int
}

type StrPriorityQueue []*strItem

func (s StrPriorityQueue) Len() int           { return len(s) }
func (s StrPriorityQueue) Less(i, j int) bool { return s[i].priority < s[j].priority }
func (s StrPriorityQueue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *StrPriorityQueue) Push(x any) {
	*s = append(*s, x.(*strItem))
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
