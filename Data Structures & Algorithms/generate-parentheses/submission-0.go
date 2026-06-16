func generateParenthesis(n int) []string {

	var result []string

	var backtrack func(open, close int, curr string)

	backtrack = func(open, close int, curr string) {

		// Found valid combination
		if len(curr) == 2*n {
			result = append(result, curr)
			return
		}

		// Add '('
		if open < n {
			backtrack(
				open+1,
				close,
				curr+"(",
			)
		}

		// Add ')'
		if close < open {
			backtrack(
				open,
				close+1,
				curr+")",
			)
		}
	}

	backtrack(0, 0, "")

	return result
}