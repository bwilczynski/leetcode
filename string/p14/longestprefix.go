package p14

func longestCommonPrefix(strs []string) string {
	prefix := []byte{}
	for i, s := range []byte(strs[0]) {
		for _, str := range strs {
			if len(str) <= i || str[i] != s {
				return string(prefix)
			}
		}
		prefix = append(prefix, s)
	}
	return string(prefix)
}
