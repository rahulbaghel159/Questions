package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3060/
func backspaceCompare(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}

	return buildString(s) == buildString(t)
}

func buildString(s string) string {
	if len(s) == 0 {
		return ""
	}

	s_str := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			s_str = append(s_str, s[i])
		} else {
			l := len(s_str)
			if l > 0 {
				s_str = s_str[:l-1]
			}
		}
	}

	return string(s_str)
}
