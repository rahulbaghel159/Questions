package graphgeneral

// https://leetcode.com/problems/number-of-islands
type DisjointSet struct {
	count  int
	parent []int
	rank   []int
}

func NewDisjointSet(grid [][]byte) *DisjointSet {
	if len(grid) == 0 {
		return &DisjointSet{}
	}

	m, n := len(grid), len(grid[0])

	d := &DisjointSet{
		count:  0,
		parent: make([]int, m*n),
		rank:   make([]int, m*n),
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				d.parent[i*n+j] = i*n + j
				d.count++
			}
			d.rank[i*n+j] = 0
		}
	}

	return d
}

func (d *DisjointSet) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}

	return d.parent[x]
}

func (d *DisjointSet) union(x, y int) {
	root_x, root_y := d.find(x), d.find(y)

	if root_x != root_y {
		rank_x, rank_y := d.rank[root_x], d.rank[root_y]

		if rank_x > rank_y {
			d.parent[root_y] = root_x
		} else if rank_y > rank_x {
			d.parent[root_x] = root_y
		} else {
			d.parent[root_y] = root_x
			d.rank[root_x]++
		}
		d.count--
	}
}

func (d *DisjointSet) getCount() int {
	return d.count
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	d := NewDisjointSet(grid)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'

				if i-1 >= 0 && grid[i-1][j] == '1' {
					d.union(i*len(grid[0])+j, (i-1)*len(grid[0])+j)
				}

				if i+1 < len(grid) && grid[i+1][j] == '1' {
					d.union(i*len(grid[0])+j, (i+1)*len(grid[0])+j)
				}

				if j-1 >= 0 && grid[i][j-1] == '1' {
					d.union(i*len(grid[0])+j, i*len(grid[0])+j-1)
				}

				if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
					d.union(i*len(grid[0])+j, i*len(grid[0])+j+1)
				}
			}
		}
	}

	return d.getCount()
}

// func numIslands(grid [][]byte) int {
// 	if grid == nil || len(grid) == 0 {
// 		return 0
// 	}

// 	num_islands, queue := 0, arrayqueue.New()
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {
// 			if grid[i][j] == '1' {
// 				num_islands++
// 				queue.Enqueue(i*len(grid[0]) + j)

// 				for queue.Size() > 0 {
// 					v, _ := queue.Dequeue()
// 					id := v.(int)

// 					r := id / len(grid[0])
// 					c := id % len(grid[0])

// 					if r-1 >= 0 && grid[r-1][c] == '1' {
// 						grid[r-1][c] = '0'
// 						queue.Enqueue((r-1)*len(grid[0]) + c)
// 					}

// 					if c-1 >= 0 && grid[r][c-1] == '1' {
// 						grid[r][c-1] = '0'
// 						queue.Enqueue(r*len(grid[0]) + c - 1)
// 					}

// 					if r+1 < len(grid) && grid[r+1][c] == '1' {
// 						grid[r+1][c] = '0'
// 						queue.Enqueue((r+1)*len(grid[0]) + c)
// 					}

// 					if c+1 < len(grid[0]) && grid[r][c+1] == '1' {
// 						grid[r][c+1] = '0'
// 						queue.Enqueue(r*len(grid[0]) + c + 1)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return num_islands
// }
