package multidp

// https://leetcode.com/problems/edit-distance/
func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	memo := make([][]int, len(word1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(word2))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	return minDistanceHelper(word1, word2, len(word1), len(word2), memo)
}

func minDistanceHelper(word1, word2 string, word1Index, word2Index int, memo [][]int) int {
	if memo[word1Index][word2Index] != -1 {
		return memo[word1Index][word2Index]
	}

	if word1Index == 0 {
		memo[word1Index][word2Index] = word2Index
		return word2Index
	}

	if word2Index == 0 {
		memo[word1Index][word2Index] = word1Index
		return word1Index
	}

	if word1[word1Index-1] == word2[word2Index-1] {
		memo[word1Index][word2Index] = minDistanceHelper(word1, word2, word1Index-1, word2Index-1, memo)
	} else {
		insert := minDistanceHelper(word1, word2, word1Index, word2Index-1, memo)
		delete := minDistanceHelper(word1, word2, word1Index-1, word2Index, memo)
		replace := minDistanceHelper(word1, word2, word1Index-1, word2Index-1, memo)

		memo[word1Index][word2Index] = min(insert, min(delete, replace)) + 1
	}

	return memo[word1Index][word2Index]
}
