package trie

// https://leetcode.com/problems/design-add-and-search-words-data-structure/
type Node struct {
	Children [26]*Node
	isLeaf   bool
}

type WordDictionary struct {
	root *Node
}

func MyConstructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, r := range word {
		if node.Children[r-'a'] == nil {
			node.Children[r-'a'] = &Node{}
		}
		node = node.Children[r-'a']
	}

	node.isLeaf = true
}

func (this *WordDictionary) Search(word string) bool {
	return find(this.root, word)
}

func find(root *Node, word string) bool {
	node := root
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			if node != nil {
				for _, n := range node.Children {
					if find(n, word[i+1:]) {
						return true
					}
				}
				return false
			} else {
				return false
			}
		} else {
			if node == nil || node.Children[word[i]-'a'] == nil {
				return false
			}
			node = node.Children[word[i]-'a']
		}
	}

	return node != nil && node.isLeaf
}
