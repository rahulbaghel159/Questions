package others

// https://leetcode.com/explore/interview/card/google/66/others-4/3102/
func numJewelsInStones(jewels string, stones string) int {
	if len(jewels) == 0 {
		return 0
	}

	dict := make(map[rune]struct{})

	for _, r := range jewels {
		dict[r] = struct{}{}
	}

	nJewels := 0
	for _, r := range stones {
		if _, ok := dict[r]; ok {
			nJewels++
		}
	}

	return nJewels
}
