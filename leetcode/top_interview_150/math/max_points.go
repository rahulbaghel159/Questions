package math

// https://leetcode.com/problems/max-points-on-a-line/
func maxPoints(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	result := 2
	for i := 0; i < len(points); i++ {
		slope_map := make(map[float64]int)
		for j := 0; j < len(points); j++ {
			if !isEqual(points[i], points[j]) {
				s := slope(points[i], points[j])
				if _, ok := slope_map[s]; ok {
					slope_map[s]++
				} else {
					slope_map[s] = 1
				}
			}
		}
		result = max(result, findMaxVal(slope_map)+1)
	}

	return result
}

func isEqual(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func slope(p1, p2 []int) float64 {
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

func findMaxVal(slope_map map[float64]int) int {
	res := 0
	for _, v := range slope_map {
		res = max(res, v)
	}

	return res
}
