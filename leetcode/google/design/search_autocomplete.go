package design

type TrieNode struct {
	children [28]*TrieNode
	dict     map[string]int
}

type AutocompleteSystem struct {
	root        *TrieNode
	word_so_far string
}

func Constructor5(sentences []string, times []int) AutocompleteSystem {
	return AutocompleteSystem{
		root:        &TrieNode{},
		word_so_far: "",
	}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	return []string{}
}

func (a *AutocompleteSystem) insertIntoTrie(s string) {
	if len(s) == 0 {
		return
	}

	node := a.root

	for _, r := range s {
		switch r {
		case ' ':
			if node.children[26] == nil {
				node.children[26] = &TrieNode{}
			}
			node = node.children[26]
		case '#':
			if node.children[27] == nil {
				node.children[27] = &TrieNode{}
			}
			node = node.children[27]
		default:
			if node.children[r-'a'] == nil {
				node.children[r-'a'] = &TrieNode{}
			}
			node = node.children[r-'a']
		}
	}
}
