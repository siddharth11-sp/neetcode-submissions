func isValidSudoku(board [][]byte) bool {
   seen := make(map[string]bool)
   for i := 0 ; i < 9 ; i++ {
    for j := 0 ; j < 9 ; j++ {
        num := board[i][j]

        if num == '.' {
            continue
        }

        row := fmt.Sprintf("%c in row %d",num, i)
        col := fmt.Sprintf("%c in col %d", num , j)
        box := fmt.Sprintf("%c in box %d-%d", num, i/3,j/3)

        if seen[row] || seen[col] || seen[box] {
            return false
        }

        seen[row] = true
        seen[col] = true
        seen[box] = true
    }
   }
   return true
}
