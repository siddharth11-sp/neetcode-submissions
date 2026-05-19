func isValid(s string) bool {
    stack := []rune{}

	pair := map[rune]rune{
		')' : '(',
		']' : '[',
		'}' : '{',
	}

	for _, ch := range s {

		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]

		if top != pair[ch]{
			return false
		}

		stack = stack[:len(stack)-1]

	}
	return len(stack) == 0

}
