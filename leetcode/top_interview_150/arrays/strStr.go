package arrays

//https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

// func strStr(haystack string, needle string) int {
// 	m, n := len(needle), len(haystack)

// 	for window_start := 0; window_start <= n-m; window_start++ {
// 		flag := true
// 		for i := 0; i <= m-1; i++ {
// 			if haystack[window_start+i] != needle[i] {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag {
// 			return window_start
// 		}
// 	}

// 	return -1
// }

// Rabin Karp Single Hash
// func strStr(haystack string, needle string) int {
// 	m, n := len(needle), len(haystack)

// 	if n < m {
// 		return -1
// 	}

// 	RADIX, MOD, MAX_WEIGHT := 26, 1000000033, 1

// 	for i := 0; i < m; i++ {
// 		MAX_WEIGHT = (MAX_WEIGHT * RADIX) % MOD
// 	}

// 	hashNeedle, hashHay := hashValue(needle, RADIX, MOD, m), 0

// 	for window_start := 0; window_start <= n-m; window_start++ {
// 		if window_start == 0 {
// 			hashHay = hashValue(haystack, RADIX, MOD, m)
// 		} else {
// 			hashHay = ((hashHay*RADIX)%MOD - ((int(haystack[window_start-1])-int('a'))*MAX_WEIGHT)%MOD + (int(haystack[window_start+m-1]) - int('a')) + MOD) % MOD
// 		}

// 		if hashNeedle == hashHay {
// 			for i := 0; i < m; i++ {
// 				if needle[i] != haystack[window_start+i] {
// 					break
// 				}
// 				if i == m-1 {
// 					return window_start
// 				}
// 			}
// 		}
// 	}

// 	return -1
// }

// func hashValue(str string, RADIX, MOD, m int) int {
// 	ans, factor := 0, 1

// 	for i := m - 1; i >= 0; i-- {
// 		ans += (int(str[i]-'a') * factor) % MOD
// 		factor = (factor * RADIX) % MOD
// 	}

// 	return ans % MOD
// }

// Rabin Karp Double Hash
// var (
// 	RADIX_1, RADIX_2, MOD_1, MOD_2 = 26, 27, 1000000033, 2147483647
// )

// func hashPair(str string, m int) []int {
// 	hash, factor1, factor2 := make([]int, 2), 1, 1

// 	for i := m - 1; i >= 0; i-- {
// 		hash[0] += (int(str[i]-'a') * factor1) % MOD_1
// 		hash[1] += (int(str[i]-'a') * factor2) % MOD_2

// 		factor1 = (factor1 * RADIX_1) % MOD_1
// 		factor2 = (factor2 * RADIX_2) % MOD_2
// 	}

// 	hash[0] = hash[0] % MOD_1
// 	hash[1] = hash[1] % MOD_2

// 	return hash
// }

// func strStr(haystack string, needle string) int {
// 	m, n := len(needle), len(haystack)

// 	if n < m {
// 		return -1
// 	}

// 	MAX_WEIGHT_1, MAX_WEIGHT_2 := 1, 1

// 	for i := 0; i < m; i++ {
// 		MAX_WEIGHT_1 = (MAX_WEIGHT_1 * RADIX_1) % MOD_1
// 		MAX_WEIGHT_2 = (MAX_WEIGHT_2 * RADIX_2) % MOD_2
// 	}

// 	hashNeedle, hashHay := hashPair(needle, m), make([]int, 2)

// 	for window_start := 0; window_start <= n-m; window_start++ {
// 		if window_start == 0 {
// 			hashHay = hashPair(haystack, m)
// 		} else {
// 			hashHay[0] = ((hashHay[0]*RADIX_1)%MOD_1 - ((int(haystack[window_start-1])-int('a'))*MAX_WEIGHT_1)%MOD_1 + (int(haystack[window_start+m-1]) - int('a')) + MOD_1) % MOD_1
// 			hashHay[1] = ((hashHay[1]*RADIX_2)%MOD_2 - ((int(haystack[window_start-1])-int('a'))*MAX_WEIGHT_2)%MOD_2 + (int(haystack[window_start+m-1]) - int('a')) + MOD_2) % MOD_2
// 		}

// 		if hashNeedle[0] == hashHay[0] && hashNeedle[1] == hashHay[1] {
// 			return window_start
// 		}
// 	}

// 	return -1
// }

// KMP
func strStr(haystack string, needle string) int {
	m, n := len(needle), len(haystack)

	if n < m {
		return -1
	}

	longest_border, haystack_pointer, needle_pointer := longestBorderArray(needle, m), 0, 0

	for haystack_pointer < n {
		if haystack[haystack_pointer] == needle[needle_pointer] {
			needle_pointer++
			haystack_pointer++

			if needle_pointer == m {
				return haystack_pointer - m
			}
		} else {
			if needle_pointer == 0 {
				haystack_pointer++
			} else {
				needle_pointer = longest_border[needle_pointer-1]
			}
		}
	}

	return -1
}

func longestBorderArray(needle string, m int) []int {
	prev, arr := 0, make([]int, m)

	arr[0] = 0

	for i := 1; i <= m-1; {
		if needle[prev] == needle[i] {
			prev++
			arr[i] = prev
			i++
		} else {
			if prev == 0 {
				arr[i] = 0
				i++
			} else {
				prev = arr[prev-1]
			}
		}
	}

	return arr
}
