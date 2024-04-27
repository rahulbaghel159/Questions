package others

import "strconv"

// https://leetcode.com/explore/interview/card/google/66/others-4/3100/
func getHint(secret string, guess string) string {
	dict_secret, dict_guess := make(map[rune]int, 0), make(map[rune]int, 0)
	bulls, cows := 0, 0

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			dict_secret[rune(secret[i])]++
			dict_guess[rune(guess[i])]++
		}
	}

	for k, v := range dict_guess {
		cows += min(v, dict_secret[k])
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
