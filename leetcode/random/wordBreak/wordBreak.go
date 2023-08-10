package wordbreak

//https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, 0)

	for _, v := range wordDict {
		wordMap[v] = true
	}

	q := queue{
		arr: []int{},
	}
	q.add(0)

	seen := make([]bool, len(s)+1)

	for !q.isEmpty() {
		start := q.remove()

		if start == len(s) {
			return true
		}

		for end := start + 1; end <= len(s); end++ {
			if seen[end] {
				continue
			}

			if _, ok := wordMap[substring(s, start, end)]; ok {
				q.add(end)
				seen[end] = true
			}
		}
	}

	return false
}

func substring(s string, start, end int) string {
	res := ""

	for i := start; i < end; i++ {
		res += string(s[i])
	}

	return res
}

type queue struct {
	arr []int
}

func (q *queue) add(n int) {
	if q.arr == nil {
		q.arr = make([]int, 0)
	}

	q.arr = append(q.arr, n)
}

func (q *queue) remove() int {
	n := q.arr[len(q.arr)-1]

	q.arr = q.arr[:len(q.arr)-1]

	return n
}

func (q *queue) isEmpty() bool {
	return len(q.arr) == 0
}
