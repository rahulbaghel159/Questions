package main

// https://leetcode.com/explore/learn/card/fun-with-arrays/523/conclusion/3231/
func thirdMax(nums []int) int {
	var m1, m2, m3 int
	var has1, has2, has3 bool

	for _, num := range nums {
		if (has1 && m1 == num) || (has2 && m2 == num) || (has3 && m3 == num) {
			continue
		}

		if !has1 || num > m1 {
			m3, has3 = m2, has2
			m2, has2 = m1, has1
			m1, has1 = num, true
		} else if !has2 || num > m2 {
			m3, has3 = m2, has2
			m2, has2 = num, true
		} else if !has3 || num > m3 {
			m3, has3 = num, true
		}
	}

	if has3 {
		return m3
	}
	return m1
}
