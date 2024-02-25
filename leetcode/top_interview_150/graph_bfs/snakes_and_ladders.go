package graphbfs

import "container/heap"

// https://leetcode.com/problems/snakes-and-ladders/
func snakesAndLadders(board [][]int) int {
	if len(board) == 0 {
		return 0
	}

	n := len(board)

	game, k, flag := make([]int, n*n+1), 1, true
	game[0] = 0
	for i := n - 1; i >= 0; i-- {
		if flag {
			for j := 0; j < n; j++ {
				game[k] = board[i][j]
				k++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				game[k] = board[i][j]
				k++
			}
		}
		flag = !flag
	}

	dist := make([]int, n*n+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = -1
	}
	dist[1] = 0
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value:    1,
		priority: 0,
	})
	for pq.Len() > 0 {
		tmp := heap.Pop(&pq).(*Item)
		node, distance := tmp.value, tmp.priority
		if distance != dist[node] {
			continue
		}
		for next := node + 1; next <= min(node+6, n*n); next++ {
			destination := next
			if game[next] != -1 {
				destination = game[next]
			}
			if dist[destination] == -1 {
				dist[destination] = dist[node] + 1
				heap.Push(&pq, &Item{
					value:    destination,
					priority: dist[destination],
				})
			}
		}
	}

	return dist[n*n]
}

// type queue struct {
// 	arr []int
// }

// func (q *queue) enqueue(x int) {
// 	q.arr = append(q.arr, x)
// }

// func (q *queue) dequeue() int {
// 	first := q.arr[0]
// 	q.arr = q.arr[1:]
// 	return first
// }

// func (q *queue) isEmpty() bool {
// 	return len(q.arr) == 0
// }

// func (q *queue) len() int {
// 	return len(q.arr)
// }

type Item struct {
	value    int
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
	item.index = len(*pq)
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
