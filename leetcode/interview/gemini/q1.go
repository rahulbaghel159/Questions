package gemini

func findIsland(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	island, visited := 0, make(map[[2]int]struct{}, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				new_island := 0
				new_island, visited = bfs(matrix, i, j, visited)
				island += new_island
			}
		}
	}

	return island
}

func bfs(matrix [][]int, r, c int, visited map[[2]int]struct{}) (int, map[[2]int]struct{}) {
	if len(matrix) == 0 {
		return 0, visited
	}

	if matrix[r][c] != 1 {
		return 0, visited
	}

	if _, ok := visited[[2]int{r, c}]; ok {
		return 0, visited
	}

	q := Queue{
		arr: make([][]int, 0),
	}
	q.Enqueue([]int{r, c})
	visited[[2]int{r, c}] = struct{}{}

	for q.Size() > 0 {
		pair := q.Dequeue()
		x, y := pair[0], pair[1]

		directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
		for _, d := range directions {
			r_new, c_new := x+d[0], y+d[1]
			if r_new >= 0 && r_new < len(matrix) && c_new >= 0 && c_new < len(matrix[0]) {
				if matrix[r_new][c_new] == 1 {
					if _, ok := visited[[2]int{r_new, c_new}]; !ok {
						q.Enqueue([]int{r_new, c_new})
					}
				}
				visited[[2]int{r_new, c_new}] = struct{}{}
			}
		}
	}

	return 1, visited
}

type Queue struct {
	arr [][]int
}

func (q *Queue) Size() int {
	if q == nil {
		return 0
	}

	return len(q.arr)
}

func (q *Queue) Enqueue(x []int) {
	if q == nil {
		return
	}

	q.arr = append(q.arr, x)
}

func (q *Queue) Dequeue() []int {
	if q == nil {
		return []int{}
	}

	n := q.Size()
	top := q.arr[n-1]
	q.arr = q.arr[:n-1]

	return top
}
