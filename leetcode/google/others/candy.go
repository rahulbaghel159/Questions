package others

// https://leetcode.com/explore/interview/card/google/66/others-4/3097/
func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	total := 0
	for i := 0; i < len(ratings); i++ {
		total = total + candies[i]
	}

	return total
}
