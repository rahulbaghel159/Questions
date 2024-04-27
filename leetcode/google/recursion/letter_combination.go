package recursion

// https://leetcode.com/explore/interview/card/google/62/recursion-4/3078/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letter_map := map[rune][]rune{
		'2': []rune("abc"),
		'3': []rune("def"),
		'4': []rune("ghi"),
		'5': []rune("jkl"),
		'6': []rune("mno"),
		'7': []rune("pqrs"),
		'8': []rune("tuv"),
		'9': []rune("wxyz"),
	}

	result := make([]string, 0)
	backtrackLetters(letter_map, []byte(""), digits, 0, &result)

	return result
}

func backtrackLetters(letter_map map[rune][]rune, curr []byte, digits string, pos int, result *[]string) {
	if len(curr) == len(digits) {
		*result = append(*result, string(curr))
		return
	}

	if pos >= len(digits) {
		return
	}

	s := letter_map[rune(digits[pos])]
	for _, r := range s {
		backtrackLetters(letter_map, append(curr, byte(r)), digits, pos+1, result)
	}
}
