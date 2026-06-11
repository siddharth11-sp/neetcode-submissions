func subsets(nums []int) [][]int {

	var result [][]int

	var dfs func(int, []int)

	dfs = func(index int, curr []int) {

		if index == len(nums) {
			temp := make([]int, len(curr))
			copy(temp, curr)
			result = append(result, temp)
			return
		}

		curr = append(curr, nums[index])
		dfs(index+1, curr)

		curr = curr[:len(curr)-1]
		dfs(index+1, curr)
	}
	dfs(0,[]int{})
	return result
}
