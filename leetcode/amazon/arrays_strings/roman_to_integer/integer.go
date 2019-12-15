package integer

//https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/2965/
func romanToInt(s string) int {
	numMap := make(map[string]int)

	numMap["I"] = 1
	numMap["IV"] = 4
	numMap["V"] = 5
	numMap["IX"] = 9
	numMap["X"] = 10
	numMap["XL"] = 40
	numMap["L"] = 50
	numMap["XC"] = 90
	numMap["C"] = 100
	numMap["CD"] = 400
	numMap["D"] = 500
	numMap["CM"] = 900
	numMap["M"] = 1000

	var ch string
	var num int
	for i := 0; i < len(s); i++ {
		ch = string(s[i])
		if (ch == "I" || ch == "X" || ch == "C") && i < len(s)-1 {
			ch += string(s[i+1])
			switch ch {
			case "I":
				num = num*10 + 1
			}
		}
	}
}
