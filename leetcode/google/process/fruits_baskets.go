package process

// https://leetcode.com/explore/interview/card/google/67/sql-2/3046/
func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}

	maxFruits := 0

	for i := 0; i < len(fruits); i++ {
		type1, type2, j := fruits[i], fruits[i], i+1
		for ; j < len(fruits); j++ {
			if fruits[j] != type1 && fruits[j] != type2 && type1 != type2 {
				maxFruits = max(maxFruits, j-i)
				break
			}

			if fruits[j] != type1 {
				type2 = fruits[j]
			}
		}

		if j == len(fruits) {
			return max(maxFruits, j-i)
		} else {
			for i+1 < len(fruits) && fruits[i] == fruits[i+1] {
				i++
			}
		}
	}

	return maxFruits
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
