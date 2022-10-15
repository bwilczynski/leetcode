package p2

func lengthOfLongestSubstring(s string) int {
	chars := make(map[byte]struct{})
	i, j := 0, 0
	max := 0
	for i < len(s)-max && j < len(s) {
		c := s[j]
		if _, ok := chars[c]; ok {
			delete(chars, s[i])
			i++
		} else {
			curr := j - i + 1
			if curr > max {
				max = curr
			}
			chars[c] = struct{}{}
			j++
		}
	}
	return max
}
