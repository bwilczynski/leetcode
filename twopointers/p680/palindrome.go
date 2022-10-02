package p680

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.
func validPalindrome(s string) bool {
	isPalindrome := func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left:right]) || isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}
