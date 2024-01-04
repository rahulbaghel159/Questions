package slidingwindow

// https://leetcode.com/problems/minimum-window-substring/

// try filtered S approach
func minWindow(s string, t string) string {
	left, right, ans, minSize := 0, 0, "", len(s)+1

	type Pair struct {
		index int
		char  rune
	}

	filteredS := make([]Pair, 0)

	count_t, count_s := make(map[rune]int), make(map[rune]int)

	for _, r := range t {
		count_t[r]++
	}

	for i, r := range s {
		if count_t[r] > 0 {
			filteredS = append(filteredS, Pair{index: i, char: r})
		}
	}

	for left <= right && right <= len(filteredS) {
		if isValidWindow(count_t, count_s) {
			if right-left < minSize {
				minSize = right - left
				ans = s[left:right]
			}
			count_s[rune(s[left])]--
			left++
			for left < len(filteredS) {
				if _, ok := count_t[rune(s[left])]; ok {
					break
				}
				count_s[rune(s[left])]--
				left++
			}
		} else {
			if right < len(s) {
				count_s[rune(s[right])]++
			}
			right++
		}
	}

	return ans
}

func isValidWindow(count_t, count_s map[rune]int) bool {
	for k, v := range count_t {
		if count_s[k] < v {
			return false
		}
	}

	return true
}
