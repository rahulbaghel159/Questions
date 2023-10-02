package arrays

//https://leetcode.com/problems/jump-game-ii

func jump(nums []int) int {
	answer, currEnd, currFar := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		currFar = max(currFar, i+nums[i])

		if i == currEnd {
			answer++
			currEnd = currFar
		}
	}

	return answer
}
