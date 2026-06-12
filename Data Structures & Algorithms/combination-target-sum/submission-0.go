func combinationSum(nums []int, target int) [][]int {
    
	var result [][]int

	var dfs func(index int, remain int, curr []int)

	dfs = func(index int, remain int, curr []int) {

		if remain == 0 {
			temp := make([]int, len(curr))
			copy(temp, curr)

			result = append(result, temp)
			return
		}
		if remain < 0 || index == len(nums) {
			return
		}
		curr = append(curr, nums[index])
		dfs(
			index, 
			remain-nums[index],
			curr,
		)
		curr = curr[:len(curr)-1]
		dfs(
			index+1,
			remain,
			curr,
		)
	}

	dfs(0, target, []int{})

	return result
}
