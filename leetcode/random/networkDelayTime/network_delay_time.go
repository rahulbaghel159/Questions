package networkdelaytime

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/network-delay-time/
var (
	MAX_INT int
	MIN_INT int
)

func networkDelayTime(times [][]int, n int, k int) int {
	MAX_INT = 999999999999999999
	MIN_INT = -999999999999999999

	if len(times) == 0 {
		return -1
	}

	adjList := make(map[int][][]int)
	for _, v := range times {
		if _, ok := adjList[v[0]]; ok {
			adjList[v[0]] = append(adjList[v[0]], []int{v[1], v[2]})
		} else {
			adjList[v[0]] = [][]int{{v[1], v[2]}}
		}
	}

	queue, visited := arrayqueue.New(), make([]int, n+1)
	for i := 1; i <= n; i++ {
		visited[i] = MAX_INT
	}

	queue.Enqueue([]int{k, 0})

	for queue.Size() > 0 {
		v, _ := queue.Dequeue()
		node := v.([]int)

		if visited[node[0]] == MAX_INT || visited[node[0]] > node[1] {
			visited[node[0]] = node[1]

			list := adjList[node[0]]
			for _, n := range list {
				if visited[n[0]] == MAX_INT || node[1]+n[1] < visited[n[0]] {
					queue.Enqueue([]int{n[0], node[1] + n[1]})
				}
			}
		}

	}

	res := MIN_INT
	for i := 1; i <= n; i++ {
		res = max(res, visited[i])
	}

	if res == MAX_INT {
		return -1
	}

	return res
}
