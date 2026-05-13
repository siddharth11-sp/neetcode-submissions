func lengthOfLongestSubstring(s string) int {
	longSubStr := 0 
	m := make(map[byte]int)

	left := 0 

	for right:= 0; right < len(s); right++{
		if idx, ok := m[s[right]]; ok && idx >= left {
			left = idx + 1 
		}

		m[s[right]] = right

		if right-left+1 > longSubStr {
			longSubStr = right-left+1
		}
	}
	return longSubStr
}
