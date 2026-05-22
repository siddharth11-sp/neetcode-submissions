
func evalRPN(tokens []string) int {
	stack := []int{}
	operator := []string{"+", "-", "*", "/"}

	for _, operand := range tokens {
		if !containsOperator(operator, operand) {
			num, _ := strconv.Atoi(operand)
			stack = append(stack, num)
			continue
		}
		op1 := stack[len(stack)-1]
		op2 := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		switch operand {
		case "+":
			stack = append(stack, op2+op1)
		case "-":
			stack = append(stack, op2-op1)
		case "*":
			stack = append(stack, op2*op1)
		case "/":
			stack = append(stack, op2/op1)
		}
	}
	return stack[0]
}

func containsOperator(list []string, target string) bool {
	for _, op := range list {
		if op == target {
			return true
		}
	}
	return false
}