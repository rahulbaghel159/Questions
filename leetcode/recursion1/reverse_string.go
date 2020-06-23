package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/250/principle-of-recursion/1440/

func reverseString(s []byte) {
	//left and right pointers
	l, r := 0, len(s)-1

	//breaking condition
	if l >= r {
		return
	}

	//processing
	var temp byte
	temp = s[l]
	s[l] = s[r]
	s[r] = temp

	//recursing
	reverseString(s[l+1 : r])
}

func reverseString2(s []byte) []byte {
	//breaking condition
	if len(s) == 1 {
		return s
	}

	//processing and recursing
	return append(reverseString2(s[1:]), s[0])
}
