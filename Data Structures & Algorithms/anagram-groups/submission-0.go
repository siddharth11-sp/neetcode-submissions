func groupAnagrams(strs []string) [][]string {
   m := make(map[string][]string)

   for _, str := range strs {
     chars := []rune(str)
     sort.Slice(chars, func(i, j int) bool{
        return chars[i] < chars[j]
     })
     m[string(chars)] = append(m[string(chars)], str)
   }

   result := [][]string{}
   for _, group := range m {
    result = append(result, group)
   }
   return result
}
