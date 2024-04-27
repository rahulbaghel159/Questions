package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3069/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	nr, nc, n_island := len(grid), len(grid[0]), 0

	queue := Queue{
		arr: make([][]int, 0),
	}

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				n_island++
				grid[r][c] = '0'

				queue.Enqueue([]int{r, c})

				for queue.Size() > 0 {
					neighbour := queue.Dequeue()
					x, y := neighbour[0], neighbour[1]

					next_arr := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

					for _, next := range next_arr {
						new_x, new_y := x+next[0], y+next[1]

						if new_x >= 0 && new_x < nr && new_y >= 0 && new_y < nc && grid[new_x][new_y] == '1' {
							grid[new_x][new_y] = '0'
							queue.Enqueue([]int{new_x, new_y})
						}
					}
				}
			}
		}
	}

	return n_island
}

type Queue struct {
	arr [][]int
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) Enqueue(x []int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Dequeue() []int {
	first := []int{-1, -1}
	if q.Size() > 0 {
		first = q.arr[0]
	}
	if q.Size() > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = [][]int{}
	}

	return first
}

func (q *Queue) Peek() []int {
	first := []int{-1, -1}
	if q.Size() > 0 {
		first = q.arr[0]
	}

	return first
}
