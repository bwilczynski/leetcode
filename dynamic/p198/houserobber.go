package p198

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max1, max2 := 0, 0
	for _, n := range nums {
		max1, max2 = max2, max(n+max1, max2)
	}
	return max2
}
