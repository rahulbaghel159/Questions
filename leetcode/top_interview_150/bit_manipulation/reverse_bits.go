package bitmanipulation

// https://leetcode.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	ret, power := uint32(0), uint32(31)
	for num != 0 {
		ret += num & 1 << power
		power -= 1
		num = num >> 1
	}

	return ret
}
