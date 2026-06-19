func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	var dfs func(index int, curr string)

	dfs = func(index int, curr string) {

		if index == len(digits) {
			result = append(result, curr)
			return
		}

		letters := mapping[digits[index]]

		for _, ch := range letters {
			dfs(index+1, curr+string(ch))
		}
	}

	dfs(0, "")

	return result
}