package feb2024

// https://leetcode.com/contest/biweekly-contest-123/problems/type-of-triangle-ii/
func triangleType(nums []int) string {
	if len(nums) < 3 {
		return "none"
	}

	s1, s2, s3 := nums[0], nums[1], nums[2]

	if s1 == s2 && s2 == s3 {
		return "equilateral"
	}

	if s1+s2 <= s3 || s2+s3 <= s1 || s1+s3 <= s2 {
		return "none"
	}

	if s1 == s2 || s2 == s3 || s1 == s3 {
		return "isosceles"
	}

	return "scalene"
}
