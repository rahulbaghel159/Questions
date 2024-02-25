package trie

// https://leetcode.com/problems/implement-trie-prefix-tree/
type TrieNode struct {
	children [26]*TrieNode
	isLeaf   bool
	word     string
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	node := this.root
	for _, r := range word {
		if node.children[r-'a'] == nil {
			node.children[r-'a'] = &TrieNode{}
		}
		node = node.children[r-'a']
	}

	node.isLeaf = true
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	node := this.root
	for _, r := range word {
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}

	return node.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	node := this.root
	for _, r := range prefix {
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}

	return true
}
