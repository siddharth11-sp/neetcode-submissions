func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    m := make(map[int]bool)
    for _, num := range nums {
        m[num]= true
    }
    lc := 1
    for _, num := range nums {
        tempLc := lcfunc(m,num+1)
        if tempLc > lc {
            lc = tempLc
        }
    }
    return lc
}

func lcfunc(m map[int]bool, num int) int {
    lc := 1
    for i := 0 ; i < len(m) ; i++ {
        if  _, ok := m[num]; ok {
            lc++
            num = num+1
        }else {
            return lc
        }
    }
    return lc
}
