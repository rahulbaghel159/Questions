package recursion

// https://leetcode.com/explore/interview/card/google/62/recursion-4/399/
func findStrobogrammatic(n int) []string {
	reversible := [][]string{{"0", "0"}, {"1", "1"}, {"6", "9"}, {"8", "8"}, {"9", "6"}}
	return recursiveFind(reversible, n, n)
}

func recursiveFind(reversible [][]string, n, final_length int) []string {
	if n == 0 {
		return []string{""}
	}

	if n == 1 {
		return []string{"0", "1", "8"}
	}

	middle := recursiveFind(reversible, n-2, n)
	result := make([]string, 0)
	for _, s := range middle {
		for _, pair := range reversible {
			if pair[0] != "0" || n != final_length {
				result = append(result, pair[0]+s+pair[1])
			}
		}
	}

	return result
}
