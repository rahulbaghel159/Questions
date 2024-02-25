package maze2

import (
	"github.com/emirpasic/gods/queues/arrayqueue"
)

// https://leetcode.com/problems/the-maze-ii
var (
	MAX_INT int
)

func shortestDistance(maze [][]int, start []int, destination []int) int {
	MAX_INT = 999999999999999999

	if len(maze) == 0 {
		return -1
	}

	distance := make([][]int, len(maze))
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(maze[0]))
		for j := 0; j < len(distance[i]); j++ {
			distance[i][j] = MAX_INT
		}
	}

	distance[start[0]][start[1]] = 0
	queue := arrayqueue.New()
	queue.Enqueue(start)
	for queue.Size() > 0 {
		v, _ := queue.Dequeue()
		currNode := v.([]int)
		directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		for _, dir := range directions {
			x, y, count := currNode[0]+dir[0], currNode[1]+dir[1], 0
			for x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0]) && maze[x][y] == 0 {
				x += dir[0]
				y += dir[1]
				count++
			}
			if x-dir[0] >= 0 && x-dir[0] < len(distance) && y-dir[1] >= 0 && y-dir[1] < len(distance[0]) && distance[currNode[0]][currNode[1]]+count < distance[x-dir[0]][y-dir[1]] {
				distance[x-dir[0]][y-dir[1]] = distance[currNode[0]][currNode[1]] + count
				queue.Enqueue([]int{x - dir[0], y - dir[1]})
			}
		}
	}

	if distance[destination[0]][destination[1]] == MAX_INT {
		return -1
	}

	return distance[destination[0]][destination[1]]
}
