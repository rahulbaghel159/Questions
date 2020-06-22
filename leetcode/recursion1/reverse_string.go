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
