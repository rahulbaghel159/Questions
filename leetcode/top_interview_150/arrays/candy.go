package arrays

//https://leetcode.com/problems/candy

func candy(ratings []int) int {
	answer := make([]int, len(ratings))

	for i := range ratings {
		answer[i] = 1
	}

	for i := range ratings {
		if i > 0 && ratings[i] > ratings[i-1] {
			answer[i] = answer[i-1] + 1
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			answer[i] = max(answer[i+1]+1, answer[i])
		}
	}

	result := 0
	for i := range answer {
		result += answer[i]
	}

	return result
}
