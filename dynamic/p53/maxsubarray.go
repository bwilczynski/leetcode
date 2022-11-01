package p53

func maxSubArray(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res, curr := nums[0], nums[0]
	for _, n := range nums[1:] {
		curr = max(n+curr, n)
		res = max(res, curr)
	}
	return res
}
