package pointers

// https://leetcode.com/problems/is-subsequence
func isSubsequence(s, t string) bool {
	pLeft, pRight, leftBound, rightBound := 0, 0, len(s), len(t)

	for pLeft < leftBound && pRight < rightBound {
		if s[pLeft] == t[pRight] {
			pLeft++
		}
		pRight++
	}

	return pLeft == leftBound
}
