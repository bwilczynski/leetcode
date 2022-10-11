package p26

func removeDuplicates(nums []int) int {
	last, k := -101, 0
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val > last {
			last = val
			nums[k] = val
			k++
		}
	}
	return k
}
