func isAnagram(s string, t string) bool {
    var alpha [26]int
    for _, ch := range s {
        pos := ch - 'a'
        alpha[pos] += 1
    }
    for _, ch := range t {
        pos := ch - 'a'
        alpha[pos] -= 1
    }

    for _, count := range alpha {
        if count > 0 || count < 0 {
            return false
        }
    }
    return true 
}
