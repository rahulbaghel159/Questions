package random

func isRobotBounded(instructions string) bool {
	var (
		directions [][]int
		x, y, face int
	)

	directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	x, y, face = 0, 0, 0

	for _, r := range instructions {
		if r == 'R' {
			face = (face + 1) % 4
		} else if r == 'L' {
			face = (face + 3) % 4
		} else if r == 'G' {
			x += directions[face][0]
			y += directions[face][1]
		}
	}

	return face != 0 || (x == 0 && y == 0)
}
