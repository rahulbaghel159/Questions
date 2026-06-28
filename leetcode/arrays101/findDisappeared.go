package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/523/conclusion/3270/
func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		idx := abs(num) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
