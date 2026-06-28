package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3247/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}

	k := 0

	for i, j := 0, len(nums)-1; i <= j; i++ {
		if nums[i] == val {
			for nums[j] == val && j > 0 {
				j--
			}
			if i >= j {
				break
			}
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
		k++
	}

	return k
}
