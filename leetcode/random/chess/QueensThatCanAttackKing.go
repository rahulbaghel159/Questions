package chess

const (
	size = 8
)

//Position denotes the co-ordinates on chess board
type Position struct {
	i int
	j int
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	queensPos := make([]Position, 0)
	for i := range queens {
		queensPos = append(queensPos, Position{i: queens[i][0], j: queens[i][1]})
	}
	kingPos := Position{i: king[0], j: king[1]}
	possibleQueens := possibleQueens(queensPos, kingPos)
	result := make([][]int, 0)
	for _, v := range possibleQueens {
		result = append(result, []int{v.i, v.j})
	}
	return result
}

func possibleQueens(queens []Position, king Position) (possibleQueens []Position) {
	moves := [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	allQueens := make(map[Position]bool)
	for _, queen := range queens {
		allQueens[queen] = true
	}
	possibleQueens = make([]Position, 0)
	for _, move := range moves {
		found, dangerousQueen := findNearestQueen(allQueens, king, move)
		if found {
			possibleQueens = append(possibleQueens, dangerousQueen)
		}
	}
	return possibleQueens
}

func findNearestQueen(allQueens map[Position]bool, king Position, move [2]int) (bool, Position) {
	found := false
	var nextMove Position
	lastPosition := king
	for {
		nextMove = Position{
			i: lastPosition.i + move[0],
			j: lastPosition.j + move[1],
		}
		if !isPossibleMove(nextMove) {
			break
		}
		if _, ok := allQueens[nextMove]; ok {
			found = true
			break
		}
		lastPosition = nextMove
	}
	return found, nextMove
}

func isPossibleMove(pos Position) bool {
	return (pos.i >= 0 && pos.i < size && pos.j >= 0 && pos.j < size)
}
