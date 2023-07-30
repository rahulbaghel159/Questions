package predictwinner

//https://leetcode.com/problems/predict-the-winner/

func PredictTheWinner(nums []int) bool {
	memo := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			memo[i][j] = -1
		}
	}

	score, _ := maxDiff(nums, 0, len(nums)-1, memo)

	return score >= 0
}

func maxDiff(nums []int, left, right int, memo [][]int) (int, [][]int) {
	if memo[left][right] != -1 {
		return memo[left][right], memo
	}

	if left == right {
		return nums[left], memo
	}

	leftScore, memo := maxDiff(nums, left+1, right, memo)
	rightScore, memo := maxDiff(nums, left, right-1, memo)

	score := max(nums[left]-leftScore, nums[right]-rightScore)

	memo[left][right] = score

	return score, memo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
