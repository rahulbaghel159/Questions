package backtracking

// https://leetcode.com/problems/combination-sum/
var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	ans = make([][]int, 0)
	combinationSumHelper(candidates, []int{}, 0, target)

	return ans
}

func combinationSumHelper(candidates []int, curr []int, currSum int, target int) {
	if currSum == target {
		ans = custom_append(ans, curr)
		return
	}

	if currSum > target {
		return
	}

	for i := 0; i < len(candidates); i++ {
		new_curr := append(curr, candidates[i])
		if len(curr) != 0 {
			last := curr[len(curr)-1]
			if candidates[i] >= last {
				combinationSumHelper(candidates, new_curr, currSum+candidates[i], target)
			}
		} else {
			combinationSumHelper(candidates, new_curr, currSum+candidates[i], target)
		}
	}
}

func custom_append(ans [][]int, curr []int) [][]int {
	new_ans := make([][]int, len(ans)+1)
	for i := 0; i < len(ans); i++ {
		new_ans[i] = ans[i]
	}

	new_ans[len(new_ans)-1] = make([]int, len(curr))
	for j := 0; j < len(curr); j++ {
		new_ans[len(new_ans)-1][j] = curr[j]
	}

	return new_ans
}
