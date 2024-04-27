package others

import "math"

// https://leetcode.com/explore/interview/card/google/66/others-4/3105/
func minAreaRect(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	points_map := make(map[int]map[int]struct{})

	for _, point := range points {
		if _, ok := points_map[point[0]]; !ok {
			points_map[point[0]] = make(map[int]struct{})
		}
		points_map[point[0]][point[1]] = struct{}{}
	}

	area := math.MaxInt
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]

			if x1 != x2 && y1 != y2 {
				_, ok1 := points_map[x1][y2]
				_, ok2 := points_map[x2][y1]

				if ok1 && ok2 {
					area = min(area, abs(x2-x1)*abs(y2-y1))
				}
			}
		}
	}

	if area == math.MaxInt {
		return 0
	}

	return area
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
