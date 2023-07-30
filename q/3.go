package main

func isEqual(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	charMap := countCharacters(str1)

	//a=2, b=1
	for _, r := range str2 {
		if _, ok := charMap[r]; ok {
			charMap[r]--
		} else {
			return false
		}
	}

	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}

	return true
}

func countCharacters(str string) map[rune]int {
	var (
		countMap map[rune]int
	)

	countMap = make(map[rune]int)

	for _, r := range str {
		if _, ok := countMap[r]; !ok {
			countMap[r] = 1
		} else {
			countMap[r]++
		}
	}

	return countMap
}

// func main() {
// 	fmt.Println(isEqual("aab", "aba"))
// 	fmt.Println(isEqual("aab", "abc"))
// 	fmt.Println(isEqual("aab", "abcd"))
// }
