package distancek

//https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := make(map[int][]int)

	graph = buildGraph(graph, root, nil)

	visited := make(map[int]bool)
	visited[target.Val] = true

	return dfs(graph, visited, []int{}, target.Val, 0, k)
}

func buildGraph(graph map[int][]int, curr, parent *TreeNode) map[int][]int {
	if curr != nil && parent != nil {
		if _, ok := graph[parent.Val]; !ok {
			graph[parent.Val] = []int{curr.Val}
		} else {
			graph[parent.Val] = append(graph[parent.Val], curr.Val)
		}

		if _, ok := graph[curr.Val]; !ok {
			graph[curr.Val] = []int{parent.Val}
		} else {
			graph[curr.Val] = append(graph[curr.Val], parent.Val)
		}
	}

	if curr.Left != nil {
		graph = buildGraph(graph, curr.Left, curr)
	}

	if curr.Right != nil {
		graph = buildGraph(graph, curr.Right, curr)
	}

	return graph
}

func dfs(graph map[int][]int, visited map[int]bool, answer []int, curr, distance, k int) []int {
	if distance == k {
		answer = append(answer, curr)
	}

	if arr, ok := graph[curr]; ok {
		for _, v := range arr {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				answer = dfs(graph, visited, answer, v, distance+1, k)
			}
		}
	}

	return answer
}
