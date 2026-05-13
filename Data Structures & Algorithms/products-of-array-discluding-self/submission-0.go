func productExceptSelf(nums []int) []int {
    output := make( []int, len(nums))
    output[0] = 1
    for i := 1; i < len(nums); i++ {
        output[i] = output[i-1] * nums[i-1]
    }

    suffix := 1
    for i := len(nums)-1; i >= 0 ; i-- {
        output[i] *= suffix
        suffix *= nums[i]
    } 

    return output
}
