package hashmap

// https://leetcode.com/problems/ransom-note
func canConstruct(ransomNote string, magazine string) bool {
	note_map := make(map[rune]int)

	for _, r := range magazine {
		note_map[r]++
	}

	for _, r := range ransomNote {
		if v, ok := note_map[r]; !ok || v == 0 {
			return false
		}
		note_map[r]--
	}

	return true
}
