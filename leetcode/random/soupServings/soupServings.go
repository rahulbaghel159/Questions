package soupservings

import "math"

//https://leetcode.com/problems/soup-servings/

func soupServings(n int) float64 {
	m := int(math.Ceil(float64(n) / 25.0))

	memo := make(map[int]map[int]float64)

	for k := 1; k <= m; k++ {
		if calculateDP(k, k, memo) > 1-1e-5 {
			return 1.0
		}
	}

	return calculateDP(m, m, memo)
}

func calculateDP(i, j int, memo map[int]map[int]float64) float64 {
	if i <= 0 && j <= 0 {
		return 0.5
	}

	if i <= 0 {
		return 1.0
	}

	if j <= 0 {
		return 0.0
	}

	if _, ok := memo[i]; ok {
		if _, ok1 := memo[i][j]; ok1 {
			return memo[i][j]
		}
	}

	result := (calculateDP(i-4, j, memo) + calculateDP(i-3, j-1, memo) + calculateDP(i-2, j-2, memo) + calculateDP(i-1, j-3, memo)) / 4.0

	if _, ok := memo[i]; !ok {
		memo[i] = make(map[int]float64)
	}
	memo[i][j] = result

	return result
}
