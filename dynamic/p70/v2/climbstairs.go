package v2

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	k, l := 2, 3
	for i := 4; i <= n; i++ {
		k, l = l, k+l
	}
	return l
}
