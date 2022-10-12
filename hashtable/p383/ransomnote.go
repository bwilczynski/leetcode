package p383

func canConstruct(ransomNote string, magazine string) bool {
	dict := func(text string) map[rune]int {
		counter := make(map[rune]int)
		for _, l := range text {
			counter[l]++
		}
		return counter
	}

	avail := dict(magazine)
	provided := dict(ransomNote)
	// check if provided contains only letters from avail
	for key, count := range provided {
		if count > avail[key] {
			return false
		}
	}
	return true
}
