package twilio

// https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/
func minSum(nums1 []int, nums2 []int) int64 {
	var s1, o1, s2, o2 int64

	for _, num := range nums1 {
		if num == 0 {
			o1++
		}
		s1 += max(int64(num), 1)
	}

	for _, num := range nums2 {
		if num == 0 {
			o2++
		}
		s2 += max(int64(num), 1)
	}

	if (s1 < s2 && o1 == 0) || (s2 < s1 && o2 == 0) {
		return -1
	}

	return max(s1, s2)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
