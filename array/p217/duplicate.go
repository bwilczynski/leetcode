package p217

func containsDuplicate(nums []int) bool {
	dict := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return true
		}
		dict[num] = struct{}{}
	}
	return false
}
