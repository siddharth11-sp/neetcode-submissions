func exist(board [][]byte, word string) bool {

	rows := len(board)
	cols := len(board[0])

	var dfs func(r, c, index int) bool

	dfs = func(r, c, index int) bool {

		// Entire word matched
		if index == len(word) {
			return true
		}

		// Out of bounds
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return false
		}

		// Character mismatch
		if board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		found :=
			dfs(r+1, c, index+1) ||
				dfs(r-1, c, index+1) ||
				dfs(r, c+1, index+1) ||
				dfs(r, c-1, index+1)

		board[r][c] = temp

		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}