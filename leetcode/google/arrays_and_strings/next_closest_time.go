package google

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/471/
func nextClosestTime(time string) string {
	h, _ := strconv.Atoi(time[:2])
	m, _ := strconv.Atoi(time[3:])

	curr := h*60 + m
	allowed := make(map[int]struct{})
	for _, r := range time {
		if r != ':' {
			allowed[int(r-'0')] = struct{}{}
		}
	}

	for {
		curr = (curr + 1) % (24 * 60)
		digits := []int{curr / 60 / 10, curr / 60 % 10, curr % 60 / 10, curr % 60 % 10}
		flag := true
		for _, d := range digits {
			if _, ok := allowed[d]; !ok {
				flag = false
			}
			if !flag {
				break
			}
		}
		if flag {
			break
		}
	}

	return fmt.Sprintf("%02d:%02d", curr/60, curr%60)
}
