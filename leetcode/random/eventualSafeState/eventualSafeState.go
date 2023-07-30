package eventualsafestate

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/find-eventual-safe-states/

func eventualSafeNodes(graph [][]int) []int {
	result, safe := make([]int, 0), make(map[int]bool)

	for i, u := range graph {
		if len(u) == 0 {
			//terminal node found
			result = append(result, i)
			safe[i] = true
		}
	}

	fmt.Println(result)

	q := &queue{arr: []int{}}

	//add all non terminal nodes to queue
	for i := 0; i < len(graph); i++ {
		if _, ok := safe[i]; !ok {
			q.add(i)
		}
	}

	for !q.isEmpty() {
		j := q.remove()
		safeFlag := 1
		for _, v := range graph[j] {
			safetyFlag, exists := safe[v]
			if !safetyFlag {
				safeFlag = 0
				break
			}
			if !exists {
				safeFlag = -1
				break
			}
		}

		switch safeFlag {
		case -1:
			q.add(j)
		case 0:
			safe[j] = false
		case 1:
			result = append(result, j)
			safe[j] = true
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

type queue struct {
	arr []int
}

func (q *queue) add(a int) {
	q.arr = append(q.arr, a)
}

func (q *queue) remove() int {
	a := q.arr[0]

	q.arr = q.arr[1:len(q.arr)]

	return a
}

func (q *queue) isEmpty() bool {
	return len(q.arr) == 0
}

func (q *queue) size() int {
	return len(q.arr)
}
