func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)

	var result [][]int

	var backtrack func(start int, curr []int)

	backtrack = func(start int, curr []int) {

		temp := make([]int, len(curr))
		copy(temp, curr)

		result = append(result, temp)

		for i := start; i < len(nums); i++ {

			// Skip duplicates
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			curr = append(curr, nums[i])

			backtrack(i+1, curr)

			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []int{})

	return result
}