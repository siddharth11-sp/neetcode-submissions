func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		currentTemp := temperatures[i]
		stack = append(stack, temperatures[i])
		for j := i + 1; j < len(temperatures); j++ {
			if currentTemp >= temperatures[j] {
				stack = append(stack, temperatures[j])
				if j == len(temperatures)-1 {
					stack = []int{}
				}
				continue
			}
			result[i] = len(stack)
			stack = []int{}
			break
		}
	}
	return result
}
