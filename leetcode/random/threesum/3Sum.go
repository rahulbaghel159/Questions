package threesum

//https://leetcode.com/problems/3sum/

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	//i start from left and j from right
	for i := 0; i < len(nums)-2; i++ {
		for j := len(nums) - 1; j > i; j-- {
			// If number i has become positive or j has become negative, then no need to proceed
			if nums[i] > 0 || nums[j] < 0 {
				break
			}
			// If current i is same as previous or j is same as previous jump to next iteration
			if i > 0 && nums[i] == nums[i-1] {
				break
			}
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}
			toFind := -(nums[i] + nums[j])
			// If target number is smaller than number i or greater than j then we must have dealt with it previously
			if toFind < nums[i] || toFind > nums[j] {
				continue
			}
			if _, ok := set[toFind]; ok {
				//if we found target number and it is neither number i or j then add it
				if toFind != nums[i] && toFind != nums[j] {
					result = append(result, []int{nums[i], toFind, nums[j]})
				} else {
					//if we found target number and it is equal to either i or j then we need to make sure it is not i or j. eg -1,-1,2 is ok, but -2, -1, 2 is not
					if i+1 != j && (toFind == nums[i+1] || toFind == nums[j-1]) {
						result = append(result, []int{nums[i], toFind, nums[j]})
					}
				}
			}
		}
	}
	return result
}
