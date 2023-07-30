package minspeedontime

import "math"

//https://leetcode.com/problems/minimum-speed-to-arrive-on-time

func minSpeedOnTime(dist []int, hour float64) int {
	left, right, minSpeed := 0, 10000000, -1

	for left <= right {
		mid := (right + left) / 2
		t := timeRequired(mid, dist)

		if t <= hour {
			minSpeed = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return minSpeed
}

func timeRequired(speed int, dist []int) float64 {
	t := float64(0)

	for i := 0; i < len(dist)-1; i++ {
		t += math.Ceil(float64(dist[i]) / float64(speed))
	}

	t += float64(dist[len(dist)-1]) / float64(speed)

	return t
}
