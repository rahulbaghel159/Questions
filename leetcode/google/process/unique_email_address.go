package process

// https://leetcode.com/explore/interview/card/google/67/sql-2/3044/
func numUniqueEmails(emails []string) int {
	if len(emails) == 0 {
		return 0
	}

	dict := make(map[string]struct{})
	for _, email := range emails {
		local, domain := splitAddress(email)
		add := rule2(rule1(local))
		if _, ok := dict[add+domain]; !ok {
			dict[add+domain] = struct{}{}
		}
	}

	return len(dict)
}

func splitAddress(add string) (string, string) {
	if len(add) == 0 {
		return "", ""
	}

	index := 0
	for i, r := range add {
		if r == '@' {
			index = i
			break
		}
	}

	return add[:index], add[index:]
}

func rule1(add string) string {
	ans := make([]rune, 0)

	for _, r := range add {
		if r != '.' {
			ans = append(ans, r)
		}
	}

	return string(ans)
}

func rule2(add string) string {
	ans := make([]rune, 0)

	for _, r := range add {
		if r != '+' {
			ans = append(ans, r)
		} else {
			break
		}
	}

	return string(ans)
}
