package adobe

// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/description/
func minOperations(s string) int {
	if len(s) < 2 {
		return 0
	}

	if len(s) == 2 {
		if s[0] != s[1] {
			return 0
		} else {
			return 1
		}
	}

	s_arr, alt1, alt2 := []byte(s), make([]byte, len(s)), make([]byte, len(s))
	for i := 0; i < len(s_arr); i++ {
		if i%2 == 0 {
			alt1[i], alt2[i] = '0', '1'
		} else {
			alt1[i], alt2[i] = '1', '0'
		}
	}

	count1, count2 := 0, 0
	for i := 0; i < len(s_arr); i++ {
		if s_arr[i] != alt1[i] {
			count1++
		}

		if s_arr[i] != alt2[i] {
			count2++
		}
	}

	return min(count1, count2)
}
