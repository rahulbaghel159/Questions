package google

import "sort"

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3062/
func kClosest(points [][]int, k int) [][]int {
	if len(points) == 0 {
		return [][]int{}
	}

	distance_arr := make([][3]int, len(points))
	for i, point := range points {
		dist := point[0]*point[0] + point[1]*point[1]
		distance_arr[i] = [3]int{dist, point[0], point[1]}
	}

	sort.Slice(distance_arr, func(i, j int) bool {
		return distance_arr[i][0] < distance_arr[j][0]
	})

	ans := make([][]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, []int{distance_arr[i][1], distance_arr[i][2]})
	}

	return ans
}
