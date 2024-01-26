package findredundantconnection

// https://leetcode.com/problems/redundant-connection/
func findRedundantConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}

	d, res := NewDisjointSet(len(edges)), make([]int, 2)
	for i := 0; i < len(edges); i++ {
		u, v := edges[i][0], edges[i][1]

		if d.find(u-1) != d.find(v-1) {
			d.union(u-1, v-1)
		} else {
			res = []int{u, v}
		}
	}

	return res
}

type DisjointSet struct {
	parent []int
	rank   []int
	n      int
}

func NewDisjointSet(n int) *DisjointSet {
	d := &DisjointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
		n:      n,
	}

	for i := 0; i < n; i++ {
		d.parent[i] = i
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
	x_root, y_root := d.find(x), d.find(y)

	if x_root == y_root {
		return
	}

	if d.rank[x_root] < d.rank[y_root] {
		d.parent[x_root] = y_root
	} else if d.rank[y_root] < d.rank[y_root] {
		d.parent[y_root] = x_root
	} else {
		d.parent[y_root] = x_root
		d.rank[x_root]++
	}
}
