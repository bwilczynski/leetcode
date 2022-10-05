package p242

func isAnagram(s string, t string) bool {
	dict := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		l := s[i]
		dict[l]++
	}
	for i := 0; i < len(t); i++ {
		l := t[i]
		if val, ok := dict[l]; ok {
			if val == 0 {
				return false
			}
			dict[l]--
			if dict[l] == 0 {
				delete(dict, l)
			}
		} else {
			return false
		}
	}
	return len(dict) == 0
}
