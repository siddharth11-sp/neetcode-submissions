func maxArea(heights []int) int {
	maxArea := 0 
	
	for i := 0; i < len(heights) - 1; i++ {
		for j := i + 1  ; j < len(heights); j++ {
			height := getHeight(heights[i], heights[j])
			area := height * (j-i)
			if area > maxArea{
				maxArea = area
			}
		}
	}

	return maxArea

}

func getHeight(h1, h2 int) int {
	if h1 > h2 {
		return h2
	}
	return h1
}
