package atlassian

// https://leetcode.com/problems/crawler-log-folder/
func minOperations(logs []string) int {
	if len(logs) == 0 {
		return 0
	}

	depth := 0
	for _, log := range logs {
		switch log {
		case "../":
			if depth > 0 {
				depth--
			}
		case "./":
		default:
			depth++
		}
	}

	return depth
}
