package atlassian

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	index := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := index[num]; ok {
			return true
		}
		index[num] = struct{}{}
	}

	return false
}
