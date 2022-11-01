package p70

var cache = make(map[int]int)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}
