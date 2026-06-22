type TrieNode struct {
	children map[byte]*TrieNode
	isWord   bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: make(map[byte]*TrieNode),
		},
	}
}

func (w *WordDictionary) AddWord(word string) {

	curr := w.root

	for i := 0; i < len(word); i++ {

		ch := word[i]

		if curr.children[ch] == nil {
			curr.children[ch] = &TrieNode{
				children: make(map[byte]*TrieNode),
			}
		}

		curr = curr.children[ch]
	}

	curr.isWord = true
}

func (w *WordDictionary) Search(word string) bool {

	var dfs func(node *TrieNode, index int) bool

	dfs = func(node *TrieNode, index int) bool {

		if index == len(word) {
			return node.isWord
		}

		ch := word[index]

		if ch == '.' {

			for _, child := range node.children {

				if dfs(child, index+1) {
					return true
				}
			}

			return false
		}

		next := node.children[ch]

		if next == nil {
			return false
		}

		return dfs(next, index+1)
	}

	return dfs(w.root, 0)
}