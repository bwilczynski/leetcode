package p28

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle) && i+j < len(haystack) && haystack[i+j] == needle[j]; j++ {
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}
