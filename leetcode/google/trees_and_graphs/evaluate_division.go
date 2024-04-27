package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/331/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	if len(queries) == 0 {
		return []float64{}
	}

	if len(equations) == 0 {
		arr := make([]float64, len(queries))
		for i := 0; i < len(arr); i++ {
			arr[i] = -1
		}
		return arr
	}

	adjList := make(map[string]map[string]float64)

	for i := 0; i < len(equations); i++ {
		dividend, divisor, cost := equations[i][0], equations[i][1], values[i]
		if _, ok := adjList[dividend]; !ok {
			adjList[dividend] = make(map[string]float64, 0)
		}
		adjList[dividend][divisor] = cost

		if _, ok := adjList[divisor]; !ok {
			adjList[divisor] = make(map[string]float64, 0)
		}
		adjList[divisor][dividend] = 1 / cost
	}

	ans := make([]float64, len(queries))
	for i, query := range queries {
		dividend, divisor := query[0], query[1]

		_, ok_dividend := adjList[dividend]
		_, ok_divisor := adjList[divisor]

		if !ok_dividend || !ok_divisor {
			ans[i] = -1
		} else {
			ans[i] = triggerDFSDivision(adjList, dividend, divisor)
		}

	}

	return ans
}

func triggerDFSDivision(adjList map[string]map[string]float64, source, destination string) float64 {
	if len(adjList) == 0 {
		return -1
	}

	if source == destination {
		return 1
	}

	visited := make(map[string]struct{})
	st := NodeStack{
		arr: make([]DivideNode, 0),
	}
	st.push(DivideNode{
		value: source,
		cost:  1,
	})
	visited[source] = struct{}{}

	for st.size() > 0 {
		node := st.pop()
		next, cost := node.value, node.cost

		if next == destination {
			return cost
		}

		for nextnode, nextcost := range adjList[next] {
			if _, ok := visited[nextnode]; !ok {
				st.push(DivideNode{
					value: nextnode,
					cost:  cost * nextcost,
				})
				visited[nextnode] = struct{}{}
			}
		}
	}

	return -1
}

type DivideNode struct {
	value string
	cost  float64
}

type NodeStack struct {
	arr []DivideNode
}

func (s *NodeStack) push(r DivideNode) {
	s.arr = append(s.arr, r)
}

func (s *NodeStack) pop() DivideNode {
	if len(s.arr) == 0 {
		return DivideNode{}
	}

	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return r
}

func (s *NodeStack) peek() DivideNode {
	if len(s.arr) == 0 {
		return DivideNode{}
	}

	return s.arr[len(s.arr)-1]
}

func (s *NodeStack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *NodeStack) size() int {
	return len(s.arr)
}
