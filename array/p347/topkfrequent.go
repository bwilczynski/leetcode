package p347

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for k, v := range freq {
		if buckets[v] == nil {
			buckets[v] = []int{k}
		} else {
			buckets[v] = append(buckets[v], k)
		}
	}

	var res []int
	for i := len(buckets) - 1; i >= 0 && k > len(res); i-- {
		if buckets[i] == nil {
			continue
		}
		for j := 0; j < len(buckets[i]) && k > len(res); j++ {
			res = append(res, buckets[i][j])
		}
	}
	return res
}
