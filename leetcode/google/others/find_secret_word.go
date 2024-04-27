package others

import "sort"

// https://leetcode.com/explore/interview/card/google/66/others-4/3104/
type Master struct{}

func (this *Master) Guess(word string) int {
	return 0
}

func findSecretWord(words []string, master *Master) {
	dict := make(map[string]int)
	//eliminated := make(map[string]struct{})

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			word1, word2 := words[i], words[j]

			if _, ok := dict[word1]; !ok {
				dict[word1] = 0
			}

			if _, ok := dict[word2]; !ok {
				dict[word2] = 0
			}

			if compare(word1, word2) > 0 {
				dict[word1] = dict[word1] + 1
				dict[word2] = dict[word1] + 1
			}
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return dict[words[i]] > dict[words[j]]
	})

	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) == 0 {
			continue
		}
		count := master.Guess(word)
		if count == 6 {
			break
		}
		cleanup(i+1, count, &words, word)
	}
}

func cleanup(start, count int, words *[]string, word string) {
	for i := start; i < len(*words); i++ {
		curr_word := (*words)[i]
		if len(curr_word) == 0 {
			continue
		}
		if compare(word, curr_word) != count {
			(*words)[i] = ""
		}
	}
}

func compare(w1, w2 string) int {
	count := 0

	for i := 0; i < len(w1); i++ {
		if w1[i] == w2[i] {
			count++
		}
	}

	return count
}
