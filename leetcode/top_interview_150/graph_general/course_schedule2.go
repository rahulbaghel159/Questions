package graphgeneral

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree, adj := make([]int, numCourses), make([][]int, numCourses)

	for i := 0; i < len(adj); i++ {
		adj[i] = make([]int, 0)
	}

	for _, prerequisite := range prerequisites {
		adj[prerequisite[1]] = append(adj[prerequisite[1]], prerequisite[0])
		inDegree[prerequisite[0]]++
	}

	visited, queue := 0, arrayqueue.New()
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.Enqueue(i)
		}
	}

	res := make([]int, 0)
	for queue.Size() > 0 {
		n, _ := queue.Dequeue()
		node := n.(int)

		visited++
		res = append(res, node)
		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue.Enqueue(neighbor)
			}
		}
	}

	if visited != numCourses {
		return []int{}
	}

	return res
}
