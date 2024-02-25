package feb2024

// https://leetcode.com/contest/biweekly-contest-123/problems/maximum-good-subarray-sum/
func maximumSubarraySum(nums []int, k int) int64 {
	if len(nums) == 0 {
		return 0
	}

	nums_map, left_sum, tmp_sum := make(map[int][]int, 0), make([]int, len(nums)), 0
	for idx, val := range nums {
		if _, ok := nums_map[val]; ok {
			nums_map[val] = append(nums_map[val], idx)
		} else {
			nums_map[val] = []int{idx}
		}
		tmp_sum += val
		left_sum[idx] = tmp_sum
	}

	var res int64
	res, tmp_sum = -999999999999999999, 0
	for i, v := range nums {
		if arr, ok := nums_map[v+k]; ok {
			for _, j := range arr {
				if i < j {
					if i != 0 {
						res = max(res, int64(left_sum[j]-left_sum[i-1]))
					} else {
						res = max(res, int64(left_sum[j]))
					}
				} else {
					if j != 0 {
						res = max(res, int64(left_sum[i]-left_sum[j-1]))
					} else {
						res = max(res, int64(left_sum[i]))
					}
				}
			}
		}
		tmp_sum += nums[i]
	}

	if res != -999999999999999999 {
		return res
	}

	return int64(0)
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
