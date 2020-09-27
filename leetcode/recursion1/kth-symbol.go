package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/253/conclusion/1675/

func kthGrammar(N int, K int) int {
	if N <= 1 {
		return 0
	}
	return (1 - K%2) ^ kthGrammar(N-1, (K+1)/2)
}
