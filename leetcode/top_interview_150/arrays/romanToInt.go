package arrays

//https://leetcode.com/problems/roman-to-integer

/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

MCDLXXVI = 1000 + 400 + 50 + 20 + 6
*/

func romanToInt(s string) int {
	value := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ans := 0

	for i := 0; i < len(s); {
		if i < len(s)-1 && value[rune(s[i])] < value[rune(s[i+1])] {
			ans += value[rune(s[i+1])] - value[rune(s[i])]
			i += 2
		} else {
			ans += value[rune(s[i])]
			i++
		}
	}

	return ans
}
