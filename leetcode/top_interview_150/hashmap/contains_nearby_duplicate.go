package hashmap

// https://leetcode.com/problems/contains-duplicate-ii
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || k == 0 {
		return false
	}

	dict := make(map[int]int)

	for i, v := range nums {
		if _, ok := dict[v]; ok {
			if i-dict[v] <= k {
				return true
			}
		}
		dict[v] = i
	}

	return false
}
