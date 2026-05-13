func isPalindrome(s string) bool {
    alphaNumStr := []rune{}
    runeList := []rune(s)
    for k := 0 ; k < len(runeList); k++{
        if isAlphaNumeric(runeList[k]){
            alphaNumStr = append(alphaNumStr,runeList[k])
        }
    }
    i := 0
    j := len(alphaNumStr) - 1
    for i < j {
        if strings.ToLower(string(alphaNumStr[i])) != strings.ToLower(string(alphaNumStr[j])) {
            return false
        }else {
            i++
            j--
        }
    }

    return true

}

func isAlphaNumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
