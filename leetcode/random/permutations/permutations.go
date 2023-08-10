package permutations

//https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	return backtrack(make([]int, 0), nums, make([][]int, 0))
}

func backtrack(curr, nums []int, ans [][]int) [][]int {
	if len(curr) == len(nums) {
		ans = append(ans, curr)
		return ans
	}

	for _, v := range nums {
		if !contains(curr, v) {
			curr = custom_append(curr, v)
			ans = backtrack(curr, nums, ans)
			curr = custom_slice(curr, len(curr)-1)
		}
	}

	return ans
}

func custom_append(arr []int, n int) []int {
	res := make([]int, len(arr)+1)

	for idx, v := range arr {
		res[idx] = v
	}

	res[len(arr)] = n

	return res
}

func custom_slice(arr []int, idx int) []int {
	if len(arr) <= 1 {
		return []int{}
	}

	res := make([]int, len(arr)-1)

	for i := 0; i < len(arr)-1; i++ {
		res[i] = arr[i]
	}

	return res
}

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if n == v {
			return true
		}
	}
	return false
}
