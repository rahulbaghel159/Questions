package arrays

//https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii

func removeDuplicates2(nums []int) int {
	j, count := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
