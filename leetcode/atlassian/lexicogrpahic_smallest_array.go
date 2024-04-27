package atlassian

import "sort"

// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/
func lexicographicallySmallestArray(nums []int, limit int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	sortedNums := make([][2]int, n)
	for i := 0; i < len(nums); i++ {
		sortedNums[i] = [2]int{nums[i], i}
	}

	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i][0] < sortedNums[j][0]
	})

	dsu := NewDisjointSet(n)
	for i := 1; i < len(nums); i++ {
		if sortedNums[i][0]-sortedNums[i-1][0] <= limit {
			dsu.union(sortedNums[i][1], sortedNums[i-1][1])
		}
	}

	subTrees := make(map[int][]int, 0)
	for i := 0; i < n; i++ {
		subTrees[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		subTrees[dsu.parent[i]] = append(subTrees[dsu.parent[i]], i)
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if dsu.parent[i] != i {
			continue
		}
		component := make([][]int, 0)
		for _, j := range subTrees[i] {
			component = append(component, []int{j, nums[j]})
		}
		sort.Slice(component, func(i, j int) bool {
			return component[i][1] < component[j][1]
		})

		for j := 0; j < len(component); j++ {
			ans[subTrees[i][j]] = component[j][1]
		}
	}

	return ans
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
