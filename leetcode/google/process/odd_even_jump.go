package process

import (
	"github.com/emirpasic/gods/maps/treemap"
)

// https://leetcode.com/explore/interview/card/google/67/sql-2/3045/
func oddEvenJumps(arr []int) int {
	canOddJump, canEvenJump := make([]bool, len(arr)), make([]bool, len(arr))

	canOddJump[len(arr)-1] = true
	canEvenJump[len(arr)-1] = true

	total := 0
	tm := treemap.NewWithIntComparator()

	for i := len(arr) - 1; i >= 0; i-- {
		if k, v := tm.Ceiling(arr[i]); k != nil {
			canOddJump[i] = canEvenJump[v.(int)]
		}

		if k, v := tm.Floor(arr[i]); k != nil {
			canEvenJump[i] = canOddJump[v.(int)]
		}

		if canOddJump[i] {
			total++
		}
		tm.Put(arr[i], i)
	}
	return total
}
