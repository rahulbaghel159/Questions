package hashmap

// https://leetcode.com/problems/two-sum
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{0}
	}

	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}

	for i, v := range nums {
		if j, ok := dict[target-v]; ok && i != j {
			return []int{i, j}
		}
	}

	return []int{0}
}
