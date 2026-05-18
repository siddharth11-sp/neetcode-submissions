func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	deque := []int{} // stores indices
	res := []int{}

	for i := 0; i < len(nums); i++ {

		// Remove indices out of current window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// Remove smaller elements
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add current index
		deque = append(deque, i)

		// Record max
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}