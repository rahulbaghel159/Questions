package roman

//https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/2964/
func intToRoman(num int) string {
	const (
		I  = 1
		IV = 4
		V  = 5
		IX = 9
		X  = 10
		XL = 40
		L  = 50
		XC = 90
		C  = 100
		CD = 400
		D  = 500
		CM = 900
		M  = 1000
	)
	str := ""
	rem := num

	for num/M > 0 && num > 0 {
		str += "M"
		rem = num % M
		num -= M
	}
	num = rem
	for num/CM > 0 && num > 0 {
		str += "CM"
		rem = num % CM
		num -= CM
	}
	num = rem
	for num/D > 0 && num > 0 {
		str += "D"
		rem = num % D
		num -= D
	}
	num = rem
	for num/CD > 0 && num > 0 {
		str += "CD"
		rem = num % CD
		num -= CD
	}
	num = rem
	for num/C > 0 && num > 0 {
		str += "C"
		rem = num % C
		num -= C
	}
	num = rem
	for num/XC > 0 && num > 0 {
		str += "XC"
		rem = num % XC
		num -= XC
	}
	num = rem
	for num/L > 0 && num > 0 {
		str += "L"
		rem = num % L
		num -= L
	}
	num = rem
	for num/XL > 0 && num > 0 {
		str += "XL"
		rem = num % XL
		num -= XL
	}
	num = rem
	for num/X > 0 && num > 0 {
		str += "X"
		rem = num % X
		num -= X
	}
	num = rem
	for num/IX > 0 && num > 0 {
		str += "IX"
		rem = num % IX
		num -= IX
	}
	num = rem
	for num/V > 0 && num > 0 {
		str += "V"
		rem = num % V
		num -= V
	}
	num = rem
	// fmt.Println("num::", num)
	for num/IV > 0 && num > 0 {
		str += "IV"
		num = 0
	}
	switch num {
	case 3:
		str += "III"
	case 2:
		str += "II"
	case 1:
		str += "I"
	}

	return str
}

//997=>500*1+100*4+9*10+5*1+1*2=>900+90+5*1+1*2=>CMXCVII
