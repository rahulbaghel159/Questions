package recursion

// https://leetcode.com/explore/interview/card/google/62/recursion-4/462/
func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	trie := &Trie{
		root: &TrieNode{},
	}
	for _, word := range words {
		trie.insert(word)
	}

	words_found := make(map[string]struct{})
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfsBoard(trie, board, i, j, string(board[i][j]), words_found)
		}
	}

	result := []string{}
	for k := range words_found {
		result = append(result, k)
	}

	return result
}

func dfsBoard(trie *Trie, board [][]byte, r, c int, prefix string, words_found map[string]struct{}) {
	if !trie.startWith(prefix) {
		return
	}

	if trie.search(prefix) {
		words_found[prefix] = struct{}{}
	}

	letter := board[r][c]
	board[r][c] = '#'

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	temp_prefix := prefix
	for _, d := range directions {
		r_new, c_new := r+d[0], c+d[1]
		if r_new >= 0 && r_new < len(board) && c_new >= 0 && c_new < len(board[0]) && board[r_new][c_new] != '#' {
			temp_prefix += string(board[r_new][c_new])
			dfsBoard(trie, board, r_new, c_new, temp_prefix, words_found)
			temp_prefix = prefix
		}
	}
	board[r][c] = letter
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	isLeaf   bool
}

func (t *Trie) insert(word string) {
	node := t.root

	for _, r := range word {
		if node.children[r-'a'] == nil {
			node.children[r-'a'] = &TrieNode{}
		}
		node = node.children[r-'a']
	}

	node.isLeaf = true
}

func (t *Trie) search(word string) bool {
	node := t.root

	for _, r := range word {
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}

	return node.isLeaf
}

func (t *Trie) startWith(prefix string) bool {
	node := t.root

	for _, r := range prefix {
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}

	return true
}
