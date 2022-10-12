package p383

func canConstruct(ransomNote string, magazine string) bool {
	avail := make(map[rune]int)
	for _, l := range magazine {
		avail[l]++
	}

	for _, l := range ransomNote {
		if avail[l] == 0 {
			return false
		}
		avail[l]--
	}
	return true
}
