package atlassian

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if len(s) == 0 || len(s) == 1 || numRows == 1 {
		return s
	}

	ans := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		j, alt_flag := i, true
		for j < len(s) {
			ans = append(ans, s[j])
			if (alt_flag || i == 0) && i != numRows-1 {
				j += (2*numRows - 2*(i+1))
			} else {
				j += 2 * i
			}
			alt_flag = !alt_flag
		}
	}

	return string(ans)
}
