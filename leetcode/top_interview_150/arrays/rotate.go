package arrays

//https://leetcode.com/problems/rotate-array

func rotate(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		current, prev := start, nums[start]
		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if current == start {
				break
			}
		}
	}
}
