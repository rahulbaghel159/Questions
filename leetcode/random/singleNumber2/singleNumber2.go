package singlenumber2

//https://leetcode.com/problems/single-number-ii

func singleNumber(nums []int) int {
	loner := int32(0)

	for shift := uint(0); shift < 32; shift++ {
		bitSum := int32(0)
		for _, num := range nums {
			bitSum += (int32(num) >> shift) & 1
		}
		lonerBit := bitSum % 3
		loner = loner | (lonerBit << shift)
	}

	return int(loner)
}
