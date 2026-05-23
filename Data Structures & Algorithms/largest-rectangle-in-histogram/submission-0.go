func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	for i := 0; i<= len(heights); i++ {

		currHeight := 0

		if i < len(heights) {
			currHeight = heights[i]
		}

		for len(stack) > 0 && currHeight < heights[stack[len(stack)-1]]{

			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i

			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := h * width

			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)

	}
	return maxArea
}
