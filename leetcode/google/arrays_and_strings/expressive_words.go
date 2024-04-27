package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3056/
func expressiveWords(s string, words []string) int {
	count := 0
	for _, w := range words {
		if isStrechy(s, w) {
			count++
		}
	}

	return count
}

func isStrechy(s, w string) bool {
	i, j := 0, 0
	for i < len(w) && j < len(s) {
		if s[j] != w[i] {
			return false
		}

		repeated1 := getRepeatedLength(w, i)
		repeated2 := getRepeatedLength(s, j)

		if ((repeated1 >= repeated2) || repeated2 < 3) && repeated1 != repeated2 {
			return false
		}

		i += repeated1
		j += repeated2
	}

	return i == len(w) && j == len(s)
}

func getRepeatedLength(word string, start int) int {
	end := start
	for end < len(word) && word[end] == word[start] {
		end++
	}
	return end - start
}
