package grpahwithshortestpath

// https://leetcode.com/problems/design-graph-with-shortest-path-calculator/
var (
	MAX_INT int
)

// type Graph struct {
// 	n       int
// 	adjList map[int][][]int
// }

// func Constructor(n int, edges [][]int) Graph {
// 	MAX_INT = 999999999999999999

// 	adjList := make(map[int][][]int)

// 	for _, tmp := range edges {
// 		if _, ok := adjList[tmp[0]]; ok {
// 			adjList[tmp[0]] = append(adjList[tmp[0]], []int{tmp[1], tmp[2]})
// 		} else {
// 			adjList[tmp[0]] = [][]int{{tmp[1], tmp[2]}}
// 		}
// 	}

// 	return Graph{
// 		n:       n,
// 		adjList: adjList,
// 	}
// }

// func (this *Graph) AddEdge(edge []int) {
// 	if _, ok := this.adjList[edge[0]]; ok {
// 		this.adjList[edge[0]] = append(this.adjList[edge[0]], []int{edge[1], edge[2]})
// 	} else {
// 		this.adjList[edge[0]] = [][]int{{edge[1], edge[2]}}
// 	}
// }

// func (this *Graph) ShortestPath(node1 int, node2 int) int {
// 	distance := make([]int, this.n)
// 	for i := 0; i < len(distance); i++ {
// 		distance[i] = MAX_INT
// 	}

// 	distance[node1] = 0

// 	pq := PriorityQueue{
// 		&Item{
// 			value:    node1,
// 			priority: 0,
// 			index:    0,
// 		},
// 	}
// 	heap.Init(&pq)

// 	for pq.Len() > 0 {
// 		tmp := heap.Pop(&pq).(*Item)
// 		cost, node := tmp.priority, tmp.value
// 		if cost > distance[node] {
// 			continue
// 		}
// 		if node == node2 {
// 			return cost
// 		}
// 		for _, neighbour := range this.adjList[node] {
// 			neighbourNode, neighbourCost := neighbour[0], neighbour[1]
// 			newcost := neighbourCost + cost
// 			if newcost < distance[neighbourNode] {
// 				distance[neighbourNode] = newcost
// 				heap.Push(&pq, &Item{
// 					value:    neighbourNode,
// 					priority: newcost,
// 				})
// 			}
// 		}
// 	}

// 	return -1
// }

// type Item struct {
// 	value    int
// 	priority int
// 	index    int
// }

// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int           { return len(pq) }
// func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// func (pq *PriorityQueue) Push(x any) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil
// 	item.index = -1
// 	*pq = old[:len(old)-1]
// 	return item
// }

type Graph struct {
	adjMatrix [][]int
}

func Constructor(n int, edges [][]int) Graph {
	MAX_INT = 999999999999999999

	adjMatrix := make([][]int, n)
	for i := 0; i < len(adjMatrix); i++ {
		adjMatrix[i] = make([]int, n)
		for j := 0; j < len(adjMatrix[i]); j++ {
			if i != j {
				adjMatrix[i][j] = MAX_INT
			} else {
				adjMatrix[i][j] = 0
			}
		}
	}

	for _, e := range edges {
		adjMatrix[e[0]][e[1]] = e[2]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				adjMatrix[j][k] = min(adjMatrix[j][k], adjMatrix[j][i]+adjMatrix[i][k])
			}
		}
	}

	return Graph{
		adjMatrix: adjMatrix,
	}
}

func (this *Graph) AddEdge(edge []int) {
	n := len(this.adjMatrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			this.adjMatrix[i][j] = min(this.adjMatrix[i][j],
				this.adjMatrix[i][edge[0]]+this.adjMatrix[edge[1]][j]+edge[2])
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	if this.adjMatrix[node1][node2] == MAX_INT {
		return -1
	}

	return this.adjMatrix[node1][node2]
}
