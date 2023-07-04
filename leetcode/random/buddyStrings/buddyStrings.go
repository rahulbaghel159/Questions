package buddystrings

//https://leetcode.com/problems/buddy-strings

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	firstIndex, secondIndex := -1, -1

	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if firstIndex == -1 {
				firstIndex = i
			} else if secondIndex == -1 {
				secondIndex = i
			} else {
				return false
			}
		}
	}

	if firstIndex == -1 && secondIndex == -1 {
		count := make([]int, 26)
		for _, v := range s {
			count[v-'a']++
			if count[v-'a'] > 1 {
				return true
			}
		}
		return false
	} else if secondIndex == -1 {
		return false
	} else if s[firstIndex] == goal[secondIndex] && s[secondIndex] == goal[firstIndex] {
		return true
	}

	return false
}
