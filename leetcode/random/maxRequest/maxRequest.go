package maxrequest

//https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/

func maximumRequests(n int, requests [][]int) int {
	return dfs(0, 0, 0, requests, make([]int, n))
}

func dfs(answer, index, count int, requests [][]int, inDegree []int) int {
	if index == len(requests) {
		for _, val := range inDegree {
			if val != 0 {
				return answer
			}
		}
		return max(answer, count)
	}

	inDegree[requests[index][0]]--
	inDegree[requests[index][1]]++

	answer = dfs(answer, index+1, count+1, requests, inDegree)

	inDegree[requests[index][0]]++
	inDegree[requests[index][1]]--

	answer = dfs(answer, index+1, count, requests, inDegree)

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
