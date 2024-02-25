package graphgeneral

// https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	//traverse top and bottom boundary
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			dfsTraverse(board, 0, j)
		}
		if board[len(board)-1][j] == 'O' {
			dfsTraverse(board, len(board)-1, j)
		}
	}

	//traverse right and left boundary
	for i := 0; i < len(board); i++ {
		if board[i][len(board[0])-1] == 'O' {
			dfsTraverse(board, i, len(board[0])-1)
		}
		if board[i][0] == 'O' {
			dfsTraverse(board, i, 0)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != 'Y' {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsTraverse(board [][]byte, i, j int) {
	q := Queue{
		arr: make([][]int, 0),
	}

	q.Enqueue([]int{i, j})
	for q.Len() > 0 {
		node := q.Dequeue()
		x, y := node[0], node[1]
		if board[x][y] != 'O' {
			continue
		}

		board[x][y] = 'Y'
		if y < len(board[0])-1 {
			q.Enqueue([]int{x, y + 1})
		}
		if x < len(board)-1 {
			q.Enqueue([]int{x + 1, y})
		}
		if y > 0 {
			q.Enqueue([]int{x, y - 1})
		}
		if x > 0 {
			q.Enqueue([]int{x - 1, y})
		}
	}
}

type Queue struct {
	arr [][]int
}

func (q *Queue) Enqueue(x []int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Dequeue() []int {
	first := q.arr[0]
	q.arr = q.arr[1:]

	return first
}

func (q *Queue) Len() int {
	return len(q.arr)
}
