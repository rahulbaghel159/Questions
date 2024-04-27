package twilio

import "sort"

// https://leetcode.com/problems/sort-array-by-increasing-frequency/
func frequencySort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	frequency_dict := make(map[int]int, 0)

	for _, num := range nums {
		frequency_dict[num]++
	}

	frequency_reverse_dict, frequency_arr := make(map[int][]int, len(frequency_dict)), make([]int, 0)
	for k, v := range frequency_dict {
		if _, ok := frequency_reverse_dict[v]; !ok {
			frequency_reverse_dict[v] = []int{}
		}
		frequency_reverse_dict[v] = append(frequency_reverse_dict[v], k)
	}

	for k := range frequency_reverse_dict {
		frequency_arr = append(frequency_arr, k)
	}

	sort.Slice(frequency_arr, func(i, j int) bool {
		return frequency_arr[i] < frequency_arr[j]
	})

	res := make([]int, 0)
	for _, frequency := range frequency_arr {
		nums := frequency_reverse_dict[frequency]
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})

		for i := 0; i < len(nums); i++ {
			for j := 0; j < frequency; j++ {
				res = append(res, nums[i])
			}
		}
	}

	return res
}
