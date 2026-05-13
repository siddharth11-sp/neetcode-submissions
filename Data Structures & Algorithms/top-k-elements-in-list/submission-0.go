func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    type Pair struct{
        Key int
        Value int
    }
    var pairs []Pair
    for _, num := range nums {
        m[num] += 1
    }

    for key, value := range m {
        pairs = append(pairs, Pair{Key : key, Value: value})
    }

    sort.Slice(pairs, func(i,j int) bool {
        return pairs[i].Value > pairs[j].Value
    } )

    result := []int{}
    for i := 0; i < k ; i++ {
        result = append(result, pairs[i].Key)
    }
    return result
}
