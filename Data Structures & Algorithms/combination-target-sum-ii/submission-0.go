func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var path []int

	var backtrack func(start, remaining int)

	backtrack = func(start, remaining int) {

		if remaining == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates); i++ {

			// Skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > remaining {
				break
			}

			path = append(path, candidates[i])

			// i+1 because element can be used only once
			backtrack(i+1, remaining-candidates[i])

			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)

	return result
}