package p1299

func replaceElements(arr []int) []int {
	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	if len(arr) == 1 {
		return arr
	}
	for i := len(arr) - 2; i >= 0; i-- {
		val := arr[i]
		arr[i] = max
		if val > max {
			max = val
		}
	}
	return arr
}
