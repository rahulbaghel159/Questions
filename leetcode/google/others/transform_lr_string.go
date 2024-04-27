package others

import (
	"strings"
)

// https://leetcode.com/explore/interview/card/google/66/others-4/3103/
func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}

	if strings.Replace(start, "X", "", len(start)) != strings.Replace(end, "X", "", len(end)) {
		return false
	}

	startL, endL, startR, endR := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)

	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			startL = append(startL, i)
		} else if start[i] == 'R' {
			startR = append(startR, i)
		}
	}

	for i := 0; i < len(end); i++ {
		if end[i] == 'L' {
			endL = append(endL, i)
		} else if end[i] == 'R' {
			endR = append(endR, i)
		}
	}

	for i := 0; i < len(startL); i++ {
		if startL[i] < endL[i] {
			return false
		}
	}

	for i := 0; i < len(startR); i++ {
		if startR[i] > endR[i] {
			return false
		}
	}

	return true
}
