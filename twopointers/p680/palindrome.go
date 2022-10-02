package p680

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			i1, j1 := i, j
			i++
			for i < j {
				if s[i] != s[j] {
					i, j = i1, j1-1
					break
				}
				i++
				j--
			}
			for i < j {
				if s[i] != s[j] {
					return false
				}
				i++
				j--
			}
			return true
		}
		i++
		j--
	}
	return true
}
