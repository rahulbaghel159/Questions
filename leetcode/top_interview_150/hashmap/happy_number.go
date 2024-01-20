package hashmap

// https://leetcode.com/problems/happy-number
func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 3 {
		return false
	}

	m, dict := n, make(map[int]bool)
	for true {
		squared_digit_sum := 0
		for m > 0 {
			d := m % 10
			squared_digit_sum += d * d
			m = m / 10
		}
		m = squared_digit_sum
		if m == 1 {
			return true
		}

		if _, ok := dict[m]; ok {
			return false
		}
		dict[m] = true
	}

	return true
}
