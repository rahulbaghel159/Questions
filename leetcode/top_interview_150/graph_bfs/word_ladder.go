package graphbfs

import "container/heap"

// https://leetcode.com/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	wordDict := make(map[string]bool, 0)
	allowed := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, v := range wordList {
		wordDict[v] = true
	}
	visited := map[string]int{
		beginWord: 0,
	}
	pq := StrPriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &StrItem{
		value:    beginWord,
		priority: 0,
	})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*StrItem)
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
					heap.Push(&pq, &StrItem{
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
