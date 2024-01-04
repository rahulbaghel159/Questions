package pointers

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.com/problems/3sum
// func threeSum(nums []int) [][]int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

// 	ans := make([][]int, 0)

// 	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
// 		if i == 0 || nums[i] != nums[i-1] {
// 			ans = append(ans, twoSum1(nums, i)...)
// 		}
// 	}

// 	return ans
// }

// func twoSum2(nums []int, i int) [][]int {
// 	res, low, high := make([][]int, 0), i+1, len(nums)-1

// 	for low < high {
// 		sum := nums[i] + nums[low] + nums[high]

// 		if sum < 0 {
// 			low++
// 		} else if sum > 0 {
// 			high--
// 		} else {
// 			res = append(res, []int{nums[i], nums[low], nums[high]})
// 			low++
// 			high--
// 			for low < high && nums[low] == nums[low-1] {
// 				low++
// 			}
// 		}
// 	}

// 	return res
// }

// func twoSum1(nums []int, i int) [][]int {
// 	seen, res := make(map[int]bool), make([][]int, 0)
// 	for j := i + 1; j < len(nums); j++ {
// 		complement := -nums[i] - nums[j]
// 		if _, ok := seen[complement]; ok {
// 			res = append(res, []int{nums[i], nums[j], complement})
// 			for j+1 < len(nums) && nums[j] == nums[j+1] {
// 				j++
// 			}
// 		}
// 		seen[nums[j]] = true
// 	}

// 	return res
// }

func threeSum(nums []int) [][]int {
	dups, seen, set := make(map[int]struct{}), make(map[int]int), make(map[string][]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dups[nums[i]]; !ok {
			dups[nums[i]] = struct{}{}
			for j := i + 1; j < len(nums); j++ {
				complement := -nums[j] - nums[i]
				if v, ok := seen[complement]; ok && v == i {
					triplet := []int{nums[i], nums[j], complement}
					sort.Slice(triplet, func(i, j int) bool {
						return triplet[i] < triplet[j]
					})
					strKey := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(triplet)), ","), "[]")
					set[strKey] = triplet
				}
				seen[nums[j]] = i
			}
		}
	}

	res, index := make([][]int, len(set)), 0

	for _, v := range set {
		res[index] = v
		index++
	}

	return res
}
