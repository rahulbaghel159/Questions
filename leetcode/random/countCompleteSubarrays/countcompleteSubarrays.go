package countcompletesubarrays

// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
func countCompleteSubarrays(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	freq, unique_count := make(map[int]int), 0
	for _, num := range nums {
		if _, ok := freq[num]; !ok {
			unique_count++
		}
		freq[num]++
	}

	return 0
}
