package cheapestflightwithkstops

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/cheapest-flights-within-k-stops/
var (
	MAX_INT, MIN_INT int
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	MAX_INT = 999999999999999999
	MIN_INT = -999999999999999999

	if len(flights) == 0 {
		return -1
	}

	adjList := make(map[int][][]int)
	for i := 0; i < len(flights); i++ {
		if _, ok := adjList[flights[i][0]]; ok {
			adjList[flights[i][0]] = append(adjList[flights[i][0]], []int{flights[i][1], flights[i][2]})
		} else {
			adjList[flights[i][0]] = [][]int{{flights[i][1], flights[i][2]}}
		}
	}

	queue, price, stops := arrayqueue.New(), make([]int, n), 0
	for i := 0; i < len(price); i++ {
		price[i] = MAX_INT
	}

	queue.Enqueue([]int{src, 0})
	for stops <= k && queue.Size() > 0 {
		sz := queue.Size()

		for sz > 0 {
			sz--
			v, _ := queue.Dequeue()
			tmp := v.([]int)
			node, distance := tmp[0], tmp[1]

			if _, ok := adjList[node]; !ok {
				continue
			}

			for _, e := range adjList[node] {
				neighbour, cost := e[0], e[1]
				if cost+distance < price[neighbour] {
					price[neighbour] = cost + distance
					queue.Enqueue([]int{neighbour, price[neighbour]})
				}
			}
		}
		stops++
	}

	if price[dst] == MAX_INT {
		return -1
	}

	return price[dst]
}
