func twoSum(nums []int, target int) []int {
    var i,j int 
    for i := 0 ; i < len(nums); i++ {
        rem := target - nums[i]
        for j := i + 1 ; j < len(nums); j++ {
            if rem == nums[j] {
                return []int{i,j }
            }
        }
    }
    return []int {i,j}
}
