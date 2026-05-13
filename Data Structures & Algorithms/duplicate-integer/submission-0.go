func hasDuplicate(nums []int) bool {
    m := make(map[int]int, len(nums))
    for _, num := range nums {
        if _, ok := m[num]; ok {
            return true
        }else {
            m[num] = 1
        }
    }
    return false
}
