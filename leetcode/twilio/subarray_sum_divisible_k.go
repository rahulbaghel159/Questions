package twilio

// https://leetcode.com/problems/subarray-sums-divisible-by-k/description/
func subarraysDivByK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	sum, count := 0, 0
	dict := make(map[int]int, 0)
	dict[0] = 1

	for _, num := range nums {
		sum += num
		rem := sum % k

		if rem < 0 {
			rem += k
		}

		count += dict[rem]
		dict[rem]++
	}

	return count
}
