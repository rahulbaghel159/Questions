package trie

// https://leetcode.com/problems/word-search-ii/
func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 {
		return []string{}
	}

	root, result := buildTrie(words), make([]string, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchPrefixInTrie(string(board[i][j]), root.root) {
				result = backtrack(board, i, j, root.root, result)
			}
		}
	}

	return result
}

func buildTrie(words []string) *Trie {
	root := &Trie{
		root: &TrieNode{},
	}
	node := root.root

	for _, word := range words {
		node = root.root
		for _, r := range word {
			if node.children[r-'a'] == nil {
				node.children[r-'a'] = &TrieNode{}
			}
			node = node.children[r-'a']
		}
		node.isLeaf = true
		node.word = word
	}

	return root
}

func searchPrefixInTrie(word string, root *TrieNode) bool {
	node := root

	for _, r := range word {
		if r == '#' {
			return false
		}
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}

	return node != nil
}

func backtrack(board [][]byte, row, col int, parent *TrieNode, result []string) []string {
	if len(board) == 0 {
		return result
	}

	letter := board[row][col]
	currNode := parent.children[letter-'a']

	if currNode.isLeaf && currNode.word != "" {
		result = append(result, currNode.word)
		currNode.word = ""
	}

	board[row][col] = '#'

	rowOffset, colOffset := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		newRow, newCol := row+rowOffset[i], col+colOffset[i]
		if newRow < 0 || newRow >= len(board) || newCol < 0 || newCol >= len(board[0]) {
			continue
		}

		if searchPrefixInTrie(string(board[newRow][newCol]), currNode) {
			result = backtrack(board, newRow, newCol, currNode, result)
		}
	}

	board[row][col] = letter

	return result
}
