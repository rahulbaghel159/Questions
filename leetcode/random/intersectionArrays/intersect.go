package intersectionarrays

// https://leetcode.com/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	count := getCounts(nums1)

	ans := make([]int, 0)
	for _, num := range nums2 {
		if v, ok := count[num]; ok && v > 0 {
			ans = append(ans, num)
			count[num]--
		}
	}

	return ans
}

func getCounts(nums []int) map[int]int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	return count
}
