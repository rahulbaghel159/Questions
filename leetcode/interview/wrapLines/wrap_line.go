package wraplines

func wrapLines(words []string, m int) []string {
	if len(words) == 0 {
		return []string{}
	}

	res, str := make([]string, 0), ""

	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(str)+len(word) <= m {
			str += word
			if len(str) < m {
				str += "-"
			} else if len(str) == m {
				res = append(res, str)
				str = ""
			}
		} else {
			if str[len(str)-1] == '-' {
				str = string([]byte(str)[:len(str)-1])
			}
			res = append(res, str)
			str = word + "-"
		}
	}

	if len(str) != 0 {
		if str[len(str)-1] == '-' {
			str = string([]byte(str)[:len(str)-1])
		}
		res = append(res, str)
	}

	return res
}
