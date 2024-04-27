package recursion

// https://leetcode.com/explore/interview/card/google/62/recursion-4/370/
func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}

	prefixHashTable := make(map[string][]string)
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			prefix := word[:i]
			if _, ok := prefixHashTable[prefix]; !ok {
				prefixHashTable[prefix] = make([]string, 0)
			}
			prefixHashTable[prefix] = append(prefixHashTable[prefix], word)
		}
	}

	res := make([][]string, 0)
	for _, word := range words {
		square := make([]string, 0)
		square = append(square, word)
		squareList := make([][]string, 0)
		backtrackWordSquare(prefixHashTable, &square, &squareList)
		if len(squareList) > 0 {
			res = append(res, squareList...)
		}
	}

	return res
}

func backtrackWordSquare(prefixHashTable map[string][]string, square *[]string, squareList *[][]string) {
	if len(*square) == len((*square)[0]) {
		temp := make([]string, len(*square))
		copy(temp, *square)
		*squareList = append(*squareList, temp)
		return
	}
	prefix := getPrefix(square)
	for _, word := range prefixHashTable[prefix] {
		*square = append(*square, word)
		backtrackWordSquare(prefixHashTable, square, squareList)
		*square = (*square)[:len(*square)-1]
	}
}

func getPrefix(square *[]string) string {
	l := len(*square)
	res := ""
	for i := 0; i < l; i++ {
		res += string((*square)[i][l])
	}
	return res
}
