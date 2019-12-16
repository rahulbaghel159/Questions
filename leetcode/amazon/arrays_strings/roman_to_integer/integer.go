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
			chNext := string(s[i+1])
			if chNext == "V" || chNext == "X" || chNext == "L" || chNext == "C" || chNext == "D" || chNext == "M" {
				ch1 := ch + chNext
				switch ch1 {
				case "IV":
					num += 4
					ch = ch1
					i++
				case "IX":
					num += 9
					ch = ch1
					i++
				case "XL":
					num += 40
					ch = ch1
					i++
				case "XC":
					num += 90
					ch = ch1
					i++
				case "CD":
					num += 400
					ch = ch1
					i++
				case "CM":
					num += 900
					ch = ch1
					i++
				}
			}
		}
		switch ch {
		case "I":
			num++
		case "V":
			num += 5
		case "X":
			num += 10
		case "L":
			num += 50
		case "C":
			num += 100
		case "D":
			num += 500
		case "M":
			num += 1000
		}
	}
	return num
}
