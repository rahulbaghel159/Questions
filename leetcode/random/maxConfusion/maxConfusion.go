package maxconfusion

//https://leetcode.com/problems/maximize-the-confusion-of-an-exam/

func maxConsecutiveAnswers(answerKey string, k int) int {
	start, countT, countF, maxLen := 0, 0, 0, 0

	for i := 0; i < len(answerKey); i++ {
		switch answerKey[i] {
		case 'T':
			countT++
		case 'F':
			countF++
		}

		conf := min(countT, countF)

		if conf <= k {
			maxLen++
		} else {w
			if answerKey[start] == 'T' {
				countT--
			} else {
				countF--
			}
			start++
		}
	}

	return maxLen
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
