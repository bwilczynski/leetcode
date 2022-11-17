package p2341

func numberOfPairs(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	var res [2]int
	for _, v := range freq {
		res[0] += v / 2
		if v%2 != 0 {
			res[1]++
		}
	}
	return res[:]
}
