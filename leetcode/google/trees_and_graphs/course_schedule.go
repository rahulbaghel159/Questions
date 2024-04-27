package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3070/
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree, adj, visitedNodes := make([]int, numCourses), make([][]int, numCourses), 0

	for i := 0; i < len(adj); i++ {
		adj[i] = make([]int, 0)
	}

	for _, edge := range prerequisites {
		inDegree[edge[0]]++
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	queue := IntQueue{
		arr: make([]int, 0),
	}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.Enqueue(i)
		}
	}

	schedule := make([]int, 0)
	for queue.Size() > 0 {
		node := queue.Dequeue()
		schedule = append(schedule, node)
		visitedNodes++

		for _, neighbour := range adj[node] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				queue.Enqueue(neighbour)
			}
		}

	}

	if visitedNodes != numCourses {
		return []int{}
	}

	return schedule
}

type IntQueue struct {
	arr []int
}

func (q *IntQueue) Size() int {
	return len(q.arr)
}

func (q *IntQueue) Enqueue(x int) {
	q.arr = append(q.arr, x)
}

func (q *IntQueue) Dequeue() int {
	first := -1
	if q.Size() > 0 {
		first = q.arr[0]
	}
	if q.Size() > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = []int{}
	}

	return first
}

func (q *IntQueue) Peek() int {
	first := -1
	if q.Size() > 0 {
		first = q.arr[0]
	}

	return first
}
