package graphgeneral

import "github.com/emirpasic/gods/queues/arrayqueue"

// https://leetcode.com/problems/course-schedule/
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree, adj := make([]int, numCourses), make([][]int, numCourses)

	for i := 0; i < len(adj); i++ {
		adj[i] = make([]int, 0)
	}

	for _, edge := range prerequisites {
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		inDegree[edge[0]]++
	}

	visited, queue := 0, arrayqueue.New()
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.Enqueue(i)
		}
	}
	for queue.Size() > 0 {
		n, _ := queue.Dequeue()
		node := n.(int)
		visited++

		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue.Enqueue(neighbor)
			}
		}
	}

	return visited == numCourses
}

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	adj := make([][]int, numCourses)

// 	for i := 0; i < len(adj); i++ {
// 		adj[i] = make([]int, 0)
// 	}

// 	for i := 0; i < len(prerequisites); i++ {
// 		adj[prerequisites[i][1]] = append(adj[prerequisites[i][1]], prerequisites[i][0])
// 	}

// 	visit, inStack := make([]bool, numCourses), make([]bool, numCourses)

// 	for i := 0; i < numCourses; i++ {
// 		if dfs(i, adj, visit, inStack) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func dfs(i int, adj [][]int, visit, inStack []bool) bool {
// 	if inStack[i] {
// 		return true
// 	}

// 	if visit[i] {
// 		return false
// 	}

// 	visit[i], inStack[i] = true, true

// 	for _, neighbour := range adj[i] {
// 		if dfs(neighbour, adj, visit, inStack) {
// 			return true
// 		}
// 	}

// 	inStack[i] = false

// 	return false
// }
