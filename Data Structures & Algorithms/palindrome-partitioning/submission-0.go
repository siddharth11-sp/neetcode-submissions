func partition(s string) [][]string {

	var result [][]string
	var curr []string

	var dfs func(start int)

	dfs = func(start int) {

		// Entire string consumed
		if start == len(s) {
			temp := make([]string, len(curr))
			copy(temp, curr)

			result = append(result, temp)
			return
		}

		// Try all substrings
		for end := start; end < len(s); end++ {

			substr := s[start : end+1]

			if !isPalindrome(substr) {
				continue
			}

			curr = append(curr, substr)

			dfs(end + 1)

			// backtrack
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)

	return result
}

func isPalindrome(s string) bool {

	left, right := 0, len(s)-1

	for left < right {

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}