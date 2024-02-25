package backtracking

// https://leetcode.com/problems/letter-combinations-of-a-phone-number
var letterMap map[rune][]string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letterMap = map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	if len(digits) == 1 {
		return letterMap[rune(digits[0])]
	}

	return letterCombinationsHelper(digits)
}

func letterCombinationsHelper(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	if len(digits) == 1 {
		return letterMap[rune(digits[0])]
	}

	arr, s_arr := letterCombinationsHelper(digits[1:]), letterMap[rune(digits[0])]
	res, k := make([]string, len(s_arr)*len(arr)), 0
	for i := 0; i < len(s_arr); i++ {
		for j := 0; j < len(arr); j++ {
			res[k] = s_arr[i] + arr[j]
			k++
		}
	}

	return res
}
