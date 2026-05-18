func maxSlidingWindow(nums []int, k int) []int {
	if k > len(nums) {
		return []int{}
	}

	i, j := 0, k-1
	var result []int
	for i = 0; j < len(nums) && i < len(nums); i++ {
		tempNums := make([]int, k)
		idx := 0
		for a := i; a <= j; a++ {
			tempNums[idx] = nums[a]
			idx++
		}
		max := max(tempNums)
		result = append(result, max)
		j++
	}

	return result
}

func max(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}