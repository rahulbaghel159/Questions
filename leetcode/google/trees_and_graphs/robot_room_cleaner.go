package treesandgraphs

//https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/1340/
/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

type Robot struct{}

func (robot *Robot) Move() bool { return false }
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean()     {}

type Set map[[2]int]struct{}

func cleanRoom(robot *Robot) {
	visited := make(Set)
	backtrack(robot, visited, 0, 0, 0)
}

func backtrack(robot *Robot, visited Set, row, col, d int) {
	visited[[2]int{row, col}] = struct{}{}

	robot.Clean()

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i := 0; i < 4; i++ {
		newD := (d + i) % 4

		new_row, new_col := row+directions[newD][0], col+directions[newD][1]

		if _, ok := visited[[2]int{new_row, new_col}]; !ok && robot.Move() {
			backtrack(robot, visited, new_row, new_col, newD)
			goBack(robot)
		}

		robot.TurnRight()
	}
}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
