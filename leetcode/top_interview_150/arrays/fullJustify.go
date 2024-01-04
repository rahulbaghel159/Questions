package arrays

import (
	"math"
)

// https://leetcode.com/problems/text-justification
func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)

	char_count, word_count, temp := 0, 0, make([]string, 0)
	for i := 0; i < len(words); i++ {
		char_count += len(words[i])
		word_count++
		if char_count+word_count-1 > maxWidth {
			ans = append(ans, justify(char_count-len(words[i]), temp, maxWidth))
			char_count = 0
			word_count = 0
			temp = make([]string, 0)
			i--
		} else {
			temp = append(temp, words[i])
		}
	}

	if len(temp) > 0 {
		ans = append(ans, justifyLastLine(temp, maxWidth))
	}

	return ans
}

func justify(char_count int, words []string, width int) string {
	if len(words) == 1 {
		return padSpaces(words[0], width)
	}

	space_available, ans, remaining_words := width-char_count, "", len(words)

	for _, w := range words {
		ans += w

		spaces_float := (float64(space_available)) / (float64(remaining_words - 1))
		spaces := int(math.Ceil(spaces_float))

		for i := 1; i <= spaces && i <= space_available; i++ {
			ans += " "
		}
		remaining_words--
		space_available -= spaces
	}

	return padSpaces(ans, width)
}

func justifyLastLine(words []string, width int) string {
	if len(words) == 1 {
		return padSpaces(words[0], width)
	}

	ans := ""

	for i := 0; i < len(words); i++ {
		ans += words[i]
		if i != len(words)-1 {
			ans += " "
		}
	}

	return padSpaces(ans, width)
}

func padSpaces(str string, width int) string {
	l := len(str)
	if l < width {
		for i := 0; i < width-l; i++ {
			str += " "
		}
	}

	return str
}
