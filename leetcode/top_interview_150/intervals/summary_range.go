package intervals

import "fmt"

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	arr, prev, idx := make([][]int, 0), nums[0], 0

	arr = append(arr, []int{nums[0], nums[0]})

	for i := 1; i < len(nums); i++ {
		if nums[i] == prev+1 {
			arr[idx][1] = nums[i]
		} else {
			arr = append(arr, []int{nums[i], nums[i]})
			idx++
		}
		prev = nums[i]
	}

	res := make([]string, 0)
	for _, v := range arr {
		if v[0] == v[1] {
			res = append(res, fmt.Sprint(v[0]))
		} else {
			res = append(res, fmt.Sprint(v[0])+"->"+fmt.Sprint(v[1]))
		}
	}

	return res
}
