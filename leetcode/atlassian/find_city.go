package atlassian

// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	INF := 99999

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
	}

	for i := 0; i < n; i++ {
		dist[i][i] = 0
	}

	for _, edge := range edges {
		i, j := edge[0], edge[1]
		dist[i][j] = edge[2]
		dist[j][i] = edge[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] && dist[i][k] != INF && dist[k][j] != INF {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	res, minCity := -1, INF
	for i := 0; i < n; i++ {
		currCity := 0
		for j := 0; j < n; j++ {
			if dist[i][j] <= distanceThreshold {
				currCity++
			}
		}
		if minCity >= currCity {
			minCity = currCity
			res = i
		}
	}

	return res
}
