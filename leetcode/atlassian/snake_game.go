package atlassian

// https://leetcode.com/problems/design-snake-game/
type SnakeGame struct {
	food               [][]int
	length, score      int
	row, col, currFood int
	snakePositions     CustomQueue
	height, width      int
}

func Constructor7(width int, height int, food [][]int) SnakeGame {
	board := make([][]int, height)
	queue := CustomQueue{
		arr: make([][]int, 0),
	}
	queue.Enqueue([]int{0, 0})
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, width)
	}
	return SnakeGame{
		height:         height,
		width:          width,
		food:           food,
		snakePositions: queue,
	}
}

func (this *SnakeGame) Move(direction string) int {
	newRow, newCol := this.row, this.col
	switch direction {
	case "U":
		newRow = this.row - 1
	case "D":
		newRow = this.row + 1
	case "L":
		newCol = this.col - 1
	case "R":
		newCol = this.col + 1
	}

	//case 1: hits wall after moving
	if newRow < 0 || newRow > this.height-1 || newCol < 0 || newCol > this.width-1 {
		return -1
	}

	//case 3: eats food => length and score increase and mark cell occupied by snake
	if this.currFood < len(this.food) && newRow == this.food[this.currFood][0] && newCol == this.food[this.currFood][1] {
		this.length++
		this.score++
		this.currFood++
		this.row, this.col = newRow, newCol

		this.snakePositions.Enqueue([]int{newRow, newCol})
	} else {
		//case 4: move snake by 1 unit in the given direction, take care of moving tail
		this.row, this.col = newRow, newCol
		this.snakePositions.Enqueue([]int{newRow, newCol})
		this.snakePositions.Dequeue()
	}

	//case 2: hits himself after moving
	for i := 0; i < len(this.snakePositions.arr)-1; i++ {
		pos := this.snakePositions.arr[i]
		if newRow == pos[0] && newCol == pos[1] {
			return -1
		}
	}

	return this.score
}

type CustomQueue struct {
	arr [][]int
}

func (q *CustomQueue) Size() int {
	return len(q.arr)
}

func (q *CustomQueue) Enqueue(x []int) {
	q.arr = append(q.arr, x)
}

func (q *CustomQueue) Dequeue() []int {
	first := []int{}
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

func (q *CustomQueue) Peek() []int {
	first := []int{}
	if q.Size() > 0 {
		first = q.arr[0]
	}

	return first
}
