package p20

func isValid(s string) bool {
	dict := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var tokens []byte
	for i := 0; i < len(s); i++ {
		val := s[i]
		if opening, ok := dict[val]; ok {
			if len(tokens) == 0 {
				return false
			}
			last := tokens[len(tokens)-1]
			if last != opening {
				return false
			}
			tokens = tokens[:len(tokens)-1]
		} else {
			tokens = append(tokens, val)
		}
	}
	return len(tokens) == 0
}
