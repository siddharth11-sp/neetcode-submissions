type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
    var res string
    for _, str := range strs {
        res += strconv.Itoa(len(str)) + "#" + str
    }
    return res
}

func (s *Solution) Decode(encoded string) []string {
    decodedString := []string{}
    for i := 0; i < len(encoded); {
        j := i
        for encoded[j] != '#' {
            j++
        }
        length, _ := strconv.Atoi(encoded[i:j])
        i = j + 1
        decodedString = append(decodedString, encoded[i:i+length])
        i += length
    }
    return decodedString
}