func solveNQueens(n int) [][]string {

	var result [][]string

	board := make([][]byte, n)

	for i := range board {
		board[i] = make([]byte, n)

		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	cols := make(map[int]bool)
	diag1 := make(map[int]bool) // r-c
	diag2 := make(map[int]bool) // r+c

	var backtrack func(row int)

	backtrack = func(row int) {

		// Found valid board
		if row == n {

			temp := make([]string, n)

			for i := 0; i < n; i++ {
				temp[i] = string(board[i])
			}

			result = append(result, temp)
			return
		}

		for col := 0; col < n; col++ {

			if cols[col] ||
				diag1[row-col] ||
				diag2[row+col] {
				continue
			}

			board[row][col] = 'Q'

			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			backtrack(row + 1)

			board[row][col] = '.'

			delete(cols, col)
			delete(diag1, row-col)
			delete(diag2, row+col)
		}
	}

	backtrack(0)

	return result
}