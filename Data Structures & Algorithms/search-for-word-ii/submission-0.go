type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {

	root := &TrieNode{
		children: make(map[byte]*TrieNode),
	}

	// Build Trie
	for _, word := range words {

		curr := root

		for i := 0; i < len(word); i++ {

			ch := word[i]

			if curr.children[ch] == nil {
				curr.children[ch] = &TrieNode{
					children: make(map[byte]*TrieNode),
				}
			}

			curr = curr.children[ch]
		}

		curr.word = word
	}

	rows := len(board)
	cols := len(board[0])

	var result []string

	var dfs func(r, c int, node *TrieNode)

	dfs = func(r, c int, node *TrieNode) {

		if r < 0 || c < 0 ||
			r >= rows || c >= cols {
			return
		}

		ch := board[r][c]

		if ch == '#' {
			return
		}

		next := node.children[ch]

		if next == nil {
			return
		}

		if next.word != "" {
			result = append(result, next.word)

			// prevent duplicates
			next.word = ""
		}

		board[r][c] = '#'

		dfs(r+1, c, next)
		dfs(r-1, c, next)
		dfs(r, c+1, next)
		dfs(r, c-1, next)

		board[r][c] = ch
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root)
		}
	}

	return result
}