package atlassian

import (
	"sort"
	"strconv"
)

// https://leetcode.com/problems/high-access-employees/
func findHighAccessEmployees(access_times [][]string) []string {
	if len(access_times) < 3 {
		return []string{}
	}

	access_index := make(map[string][]string)
	for _, access := range access_times {
		if _, ok := access_index[access[0]]; !ok {
			access_index[access[0]] = make([]string, 0)
		}
		access_index[access[0]] = append(access_index[access[0]], access[1])
	}

	ans := make([]string, 0)
	for emp, times := range access_index {
		if len(times) < 3 {
			continue
		}
		sort.Slice(times, func(i, j int) bool {
			return times[i] < times[j]
		})

		flag := false
		for i := 0; i+2 < len(times); i++ {
			if isWithinOneHour(times[i], times[i+2]) {
				flag = true
			}

			if flag {
				break
			}
		}

		if flag {
			ans = append(ans, emp)
		}
	}

	return ans
}

func isWithinOneHour(t1, t2 string) bool {
	h1, _ := strconv.Atoi(t1[:2])
	h2, _ := strconv.Atoi(t2[:2])
	m1, _ := strconv.Atoi(t1[2:])
	m2, _ := strconv.Atoi(t2[2:])
	diff := 0

	if h2-h1 > 1 {
		return false
	} else if h2-h1 == 1 {
		diff = (60 - m1) + m2
	} else {
		diff = m2 - m1
	}

	return diff < 60
}
