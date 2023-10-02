package arrays

//https://leetcode.com/problems/remove-duplicates-from-sorted-array

func removeDuplicates(nums []int) int {
	insertIndex := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}

	return insertIndex
}
