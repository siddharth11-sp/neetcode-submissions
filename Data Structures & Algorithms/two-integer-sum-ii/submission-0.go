func twoSum(numbers []int, target int) []int {
    i := 0 
    j := len(numbers) - 1
    result := []int{}
    for i < j {
        if numbers[i] + numbers[j] == target {
            result = append(result, i+1)
            result = append(result, j+1)
            return result
        }else if numbers[i] + numbers[j] < target {
            i++
        } else {
            j--
        }
    }
    return result
}

 