package bitmanipulation

// https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	n, m := len(a), len(b)
	if n < m {
		return addBinary(b, a)
	}

	if len(b) == 0 {
		return a
	}

	res, carry, j := "", 0, m-1
	for i := n - 1; i >= 0; i-- {
		if a[i] == '1' {
			carry++
		}
		if j >= 0 && b[j] == '1' {
			carry++
		}
		j--

		if carry%2 == 1 {
			res = "1" + res
		} else {
			res = "0" + res
		}

		carry /= 2
	}

	if carry == 1 {
		res = "1" + res
	}

	return res
}
