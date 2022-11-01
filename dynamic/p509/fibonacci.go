package p509

func fib(n int) int {
	if n < 2 {
		return n
	}
	k, l := 0, 1
	for i := 2; i <= n; i++ {
		k, l = l, k+l
	}
	return l
}
