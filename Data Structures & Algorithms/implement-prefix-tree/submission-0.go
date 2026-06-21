type trieNode struct {
	children map[byte]*trieNode
	isWord bool
}

type PrefixTree struct {
	root *trieNode
}

func Constructor() PrefixTree {
    return PrefixTree {
		root : &trieNode{
			children : make(map[byte]*trieNode),
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.root

	for i := 0 ; i < len(word); i++ {
		ch := word[i]
		if curr.children[ch] == nil {
			curr.children[ch] = &trieNode{
			children : make(map[byte]*trieNode),
		}
		}
		curr = curr.children[ch]
	}
	curr.isWord = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.root

	for i := 0 ; i < len(word); i++ {
		ch := word[i]
		if curr.children[ch] == nil {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.root

	for i := 0 ; i < len(prefix); i++ {
		ch := prefix[i]
		if curr.children[ch] == nil {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}
