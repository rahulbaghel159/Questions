package container

//https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/2963/
func maxArea(height []int) int {
	l, r, max := 0, len(height)-1, 0
	for l < r {
		temp := minNumber(height[l], height[r]) * (r - l)
		if temp > max {
			max = temp
		}
		// fmt.Printf("height[l]:%d,height[r]:%d,max:%d\n", height[l], height[r], max)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}

func minNumber(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}
